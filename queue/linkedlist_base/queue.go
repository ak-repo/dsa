package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type Queue struct {
	Front *Node
	Rear  *Node
	Size  int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) IsEmpty() bool {
	return q.Front == nil
}

func (q *Queue) Enqueue(val int) {
	newNode := &Node{Value: val}
	if q.IsEmpty() {
		q.Front = newNode
		q.Rear = newNode
	} else {
		q.Rear.Next = newNode
		q.Rear = newNode
	}
	q.Size++
}

func (q *Queue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	val := q.Front.Value
	q.Front = q.Front.Next
	if q.Front == nil {
		q.Rear = nil
	}
	q.Size--
	return val, true
}

func (q *Queue) Peek() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	return q.Front.Value, true
}

func (q *Queue) Display() {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return
	}
	current := q.Front
	fmt.Print("Front <- ")
	for current != nil {
		fmt.Printf("%d ", current.Value)
		current = current.Next
	}
	fmt.Println("<- Rear")
}

func main() {
	q := NewQueue()
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	q.Enqueue(40)
	q.Display()

	val, ok := q.Dequeue()
	if ok {
		fmt.Println("Dequeued:", val)
	}
	q.Display()
}
