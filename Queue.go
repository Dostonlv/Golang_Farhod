package main

import "fmt"

func main() {
	testQueue()
}

func testQueue() {
	queue := NewQueue()
	queue.Enqueue(100)
	fmt.Println("queue: ", queue)
	queue.Enqueue(25).Enqueue(77)
	fmt.Println("queue: ", queue)
	fmt.Println("is queue empty? ", queue.isEmpty())

	var result, _ = queue.Peek()
	fmt.Println("the next item in the queue: ", result)

	fmt.Println("Dequeue one item...")
	result, _ = queue.Dequeue()
	fmt.Println(result)
	fmt.Println("Dequeue one item...")
	result, _ = queue.Dequeue()
	fmt.Println(result)
	fmt.Println("Dequeue one item...")
	result, _ = queue.Dequeue()
	fmt.Println(result)
	fmt.Println(queue.isEmpty())
	_, err := queue.Peek()
	fmt.Println(err)
}

type Queue struct {
	data []int
}

func NewQueue() *Queue {
	return &Queue{
		data: []int{},
	}
}

func (q *Queue) isEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) Peek() (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("Queue is empty")
	}
	return q.data[0], nil
}

func (q *Queue) Enqueue(n int) *Queue {
	q.data = append(q.data, n)
	return q
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("Queue is empty")
	}
	element := q.data[0]
	q.data = q.data[1:]
	return element, nil
}
