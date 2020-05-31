package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	store []int
	front int
	rear int
	size int
}
func NewQueue(size int) *Queue{
	return &Queue{make([]int, size), 0, 0, 0}
}

func (q *Queue) isFull() bool {
	return len(q.store) == q.size
}
func (q *Queue) enque(num int) error{
	if q.size == len(q.store) {
		return errors.New("Queue full")
	}
	q.store[q.front] = num
	q.size += 1
	q.front = (q.front + 1) % len(q.store)
	return nil
}

func (q *Queue) deque() int {
	ele := q.store[q.rear]
	q.size -= 1
	q.rear = (q.rear + 1) % len(q.store)
	return ele
}

type Window struct {
	win *Queue
	avg float64
}

func (x *Window) Next(num int) {
	if x.win.isFull() {
		oldNum := x.win.deque()
		_ = x.win.enque(num)
		x.avg = (x.avg * float64(x.win.size) - float64(oldNum) + float64(num))/float64(x.win.size)
	} else {
		_ = x.win.enque(num)
		x.avg = (x.avg * float64(x.win.size - 1) + float64(num))/float64(x.win.size)
	}
}

func main() {
	w := Window{NewQueue(3),0.0}
	w.Next(5)
	w.Next(10)
	w.Next(15)
	w.Next(20)
	fmt.Println(w.avg)
}