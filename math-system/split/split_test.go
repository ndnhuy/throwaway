package split

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type someTask struct {
	x int
}

func (s someTask) execute() int {
	time.Sleep(time.Duration(1) * time.Second)
	return s.x
}

func doSumTask(inputSize int, nWorkers int) int {
	source := make(chan task[int])
	dests := Split(source, nWorkers)

	go func() {
		defer close(source)
		for i := 0; i < inputSize; i++ {
			source <- someTask{
				x: i,
			}
		}
	}()

	var wg sync.WaitGroup
	wg.Add(len(dests))

	sum := 0
	for _, d := range dests {
		go func(ch <-chan int) {
			defer wg.Done()
			for v := range ch {
				sum += v
			}
		}(d)
	}
	return sum
}

func TestSplit(t *testing.T) {
	source := make(chan task[int])
	nWorkers := 100
	dests := Split(source, nWorkers)
	require.Equal(t, nWorkers, len(dests))

	expectSum := 0
	taskSize := 2
	go func() {
		defer close(source)
		for i := 0; i < taskSize; i++ {
			source <- someTask{
				x: i,
			}
			expectSum += i
		}
	}()

	var wg sync.WaitGroup
	wg.Add(len(dests))

	sum := 0
	for _, d := range dests {
		go func(ch <-chan int) {
			defer wg.Done()
			for v := range ch {
				sum += v
			}
		}(d)
	}

	wg.Wait()

	require.Equal(t, expectSum, sum)
}
