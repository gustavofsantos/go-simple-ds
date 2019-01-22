package main

import "fmt"

// Queue Struct to handle elements
type Queue struct {
	items []int
	size  int
}

// CreateQueue Create a instance of the Queue
func CreateQueue() Queue {
	queue := Queue{size: 0}
	return queue
}

// Enqueue Enqueue a item in the Queue
func (q *Queue) Enqueue(item int) int {
	q.items = append(q.items, item)
	q.size++

	return q.size
}

// Dequeue Remove que first item in the Queue and returns it
func (q *Queue) Dequeue() int {
	item := q.items[0]

	q.items = q.items[1:]
	q.size--

	return item
}

// IsEmpty Return a boolean value about the emptyness of the queue
func (q *Queue) IsEmpty() bool {
	if q.size > 0 {
		return false
	}

	return true
}

// Length Returns the number of items in the Queue
func (q *Queue) Length() int {
	return q.size
}

// PrettyPrint Print the Queue
func (q *Queue) PrettyPrint() {
	if q.size > 7 {
		fmt.Printf("#%d - %v ...\n", q.size, q.items[:7])
	} else {
		fmt.Printf("#%d - %v\n", q.size, q.items)
	}
}
