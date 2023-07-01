package main

import (
	"fmt"
	"time"
)

func main() {
	timerChan := make(chan time.Time)
	go func() {
		time.Sleep(3 * time.Second)
		timerChan <- time.Now()
	}()
	completedAt := <-timerChan
	fmt.Println(completedAt)
}
