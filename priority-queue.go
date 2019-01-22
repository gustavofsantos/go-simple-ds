package main

import "fmt"

// PriorityQueue Simple double queue
type PriorityQueue struct {
	priority Queue
	normal   Queue
}

// CreatePriorityQueue Create a instance of the PriorityQueue
func CreatePriorityQueue() PriorityQueue {
	queue := PriorityQueue{
		priority: Queue{size: 0},
		normal:   Queue{size: 0},
	}

	return queue
}

// Enqueue Enqueue a item in the Queue
func (q *PriorityQueue) Enqueue(item int, priority bool) int {
	if priority {
		q.priority.Enqueue(item)
	} else {
		q.normal.Enqueue(item)
	}

	return q.priority.size + q.normal.size
}

// Dequeue Remove que first item in the Queue and returns it
func (q *PriorityQueue) Dequeue() int {
	if q.priority.size > 0 {
		return q.priority.Dequeue()
	}

	return q.normal.Dequeue()
}

// IsEmpty Return a boolean value about the emptyness of the queue
func (q *PriorityQueue) IsEmpty() bool {
	return q.priority.IsEmpty() && q.normal.IsEmpty()
}

// Length Returns the number of items in the Queue
func (q *PriorityQueue) Length() int {
	return q.priority.Length() + q.normal.Length()
}

// PrettyPrint Print the Queue
func (q *PriorityQueue) PrettyPrint() {
	fmt.Println("Priority Queue")
	q.priority.PrettyPrint()

	fmt.Println("Normal Queue")
	q.normal.PrettyPrint()
}
