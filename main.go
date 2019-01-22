package main

func main() {
	queue := CreatePriorityQueue()

	queue.Enqueue(1, false)
	queue.Enqueue(2, true)
	queue.Enqueue(3, false)
	queue.Enqueue(4, true)
	queue.Enqueue(5, false)

	queue.PrettyPrint()

	queue.Dequeue()

	queue.PrettyPrint()
}
