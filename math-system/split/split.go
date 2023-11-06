package split

import "sync"

type Task[R any] interface {
	Execute() R
}

func Split[T any](source <-chan Task[T], n int) []<-chan T {
	dests := make([]<-chan T, 0)
	for i := 0; i < n; i++ {
		ch := make(chan T)
		dests = append(dests, ch)
		go func() {
			defer close(ch)
			for task := range source {
				v := task.Execute()
				ch <- v
			}
		}()
	}
	return dests
}

func SplitWithOneDest[T any](source <-chan Task[T], n int) <-chan T {
	dest := make(chan T)
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for task := range source {
				v := task.Execute()
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
