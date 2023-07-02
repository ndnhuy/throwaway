package main

import (
	"fmt"
	"time"
)

type Work struct {
	x, y, z int
}

func worker(in <-chan *Work, out chan<- *Work) {
	for w := range in {
		w.z = w.x * w.y
		time.Sleep(time.Duration(1) * time.Second)
		out <- w
	}
}

func Run() {
	in, out := make(chan *Work), make(chan *Work)
	for i := 0; i < 100; i++ {
		go worker(in, out)
	}
	// send work
	go func() {
		for i := 0; i < 10; i++ {
			in <- &Work{
				x: i,
				y: i + 1,
			}
		}
	}()

	// result
	for i := 0; i < 10; i++ {
		w := <-out
		fmt.Printf("%v x %v = %v\n", w.x, w.y, w.z)
	}
	close(in)
	close(out)
}
