package main

import "fmt"

func main() {
	q := NewQueue()
	q.Put(1)
	q.Put(4)
	q.Put(2)
	q.Put(3)

	for !q.IsEmpty() {
		fmt.Println(q.Pop())
	}
}
