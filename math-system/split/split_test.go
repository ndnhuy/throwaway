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

	var mu1, mu2 sync.Mutex
	sum1 := 0
	sum2 := 0
	for i, d := range dests {
		go func(i int, ch <-chan int) {
			defer wg.Done()
			for v := range ch {
				if i%2 == 0 {
					mu1.Lock()
					sum1 += v
					mu1.Unlock()
				} else {
					mu2.Lock()
					sum2 += v
					mu2.Unlock()
				}
			}
		}(i, d)
	}

	wg.Wait()
	return sum1 + sum2
}

func doSumTaskWithOneDestChannel(inputSize int, nWorkers int, getSleepTime func() int) int {
	source := make(chan task[int])
	dest := SplitWithOneDest(source, nWorkers)

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
	sum := doSumTaskWithOneDestChannel(taskSize, 100, func() int { return 100 })
	require.Equal(t, expectSum, sum)
}

var table = []struct {
	inputSize int
	nWorkers  int
}{
	{inputSize: 10000, nWorkers: 100},
	{inputSize: 10000, nWorkers: 200},
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
				doSumTaskWithOneDestChannel(v.inputSize, v.nWorkers, func() int { return sleepTime })
			}
		})
	}
}
