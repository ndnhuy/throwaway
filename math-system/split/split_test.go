package split

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type someTask struct {
	x         int
	sleepTime int
}

func (s someTask) execute() int {
	time.Sleep(time.Duration(s.sleepTime) * time.Millisecond)
	return s.x
}

func doSumTask(inputSize int, nWorkers int, getSleepTime func() int) int {
	source := make(chan task[int])
	dests := Split(source, nWorkers)

	go func() {
		defer close(source)
		sleepTime := getSleepTime()
		for i := 0; i < inputSize; i++ {
			source <- someTask{
				x:         i,
				sleepTime: sleepTime,
			}
		}
	}()

	var wg sync.WaitGroup
	wg.Add(len(dests))

	var mu sync.Mutex
	sum := 0
	for _, d := range dests {
		go func(ch <-chan int) {
			defer wg.Done()
			for v := range ch {
				mu.Lock()
				sum += v
				mu.Unlock()
			}
		}(d)
	}

	wg.Wait()
	return sum
}

func doSumTask2(inputSize int, nWorkers int, getSleepTime func() int) int {
	source := make(chan task[int])
	dest := Split2(source, nWorkers)

	go func() {
		defer close(source)
		sleepTime := getSleepTime()
		for i := 0; i < inputSize; i++ {
			source <- someTask{
				x:         i,
				sleepTime: sleepTime,
			}
		}
	}()

	sum := 0
	for v := range dest {
		sum += v
	}

	return sum
}

func TestSplit(t *testing.T) {
	expectSum := 0
	taskSize := 100
	for i := 0; i < taskSize; i++ {
		expectSum += i
	}
	sum := doSumTask(taskSize, 100, func() int { return 100 })
	require.Equal(t, expectSum, sum)
}
func TestSplit2(t *testing.T) {
	expectSum := 0
	taskSize := 100
	for i := 0; i < taskSize; i++ {
		expectSum += i
	}
	sum := doSumTask2(taskSize, 100, func() int { return 100 })
	require.Equal(t, expectSum, sum)
}

var table = []struct {
	inputSize int
	nWorkers  int
}{
	{inputSize: 1000, nWorkers: 100},
	{inputSize: 1000, nWorkers: 200},
}

var sleepTime = 100
func BenchmarkSplit(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%v_with_%v_workers", v.inputSize, v.nWorkers), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				doSumTask(v.inputSize, v.nWorkers, func() int { return sleepTime })
			}
		})
	}
}

func BenchmarkSplit2(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%v_with_%v_workers", v.inputSize, v.nWorkers), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				doSumTask2(v.inputSize, v.nWorkers, func() int { return sleepTime })
			}
		})
	}
}
