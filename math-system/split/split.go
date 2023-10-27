package split

import "sync"

type task[R any] interface {
	execute() R
}

func Split[T any](source <-chan task[T], n int) []<-chan T {
	dests := make([]<-chan T, 0)
	for i := 0; i < n; i++ {
		ch := make(chan T)
		dests = append(dests, ch)
		go func() {
			defer close(ch)
			for task := range source {
				v := task.execute()
				ch <- v
			}
		}()
	}
	return dests
}

func SplitWithOneDest[T any](source <-chan task[T], n int) <-chan T {
	dest := make(chan T)
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for task := range source {
				v := task.execute()
				dest <- v
			}
		}()
	}
	go func() {
		wg.Wait()
		close(dest)
	}()
	return dest
}
