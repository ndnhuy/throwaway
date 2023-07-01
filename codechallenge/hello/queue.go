package main

type Queue struct {
	arr []any
}

func NewQueue() *Queue {
	return &Queue{
		arr: []any{},
	}
}

func (q *Queue) Put(e any) any {
	q.arr = append(q.arr, e)
	return e
}

func (q *Queue) Pop() any {
	rs := q.arr[0]
	q.arr = q.arr[1:]
	return rs
}

func (q *Queue) IsEmpty() bool {
	return len(q.arr) == 0
}

func (q *Queue) Size() int {
	return len(q.arr)
}
