package main

import "fmt"

// Node represents a single node in the linked list
type Node struct {
	Value int
	Next  *Node
}

// SingleList represents the singly linked list
type SingleList struct {
	Head *Node
}

// NewSingleList creates and returns a new empty linked list
func NewSingleList() *SingleList {
	return &SingleList{}
}

// IsEmpty checks if the list is empty
func (s *SingleList) IsEmpty() bool {
	return s.Head == nil
}

// Append inserts a node at the end of the list
func (s *SingleList) Append(val int) {
	newNode := &Node{Value: val}

	if s.IsEmpty() {
		s.Head = newNode
		return
	}

	current := s.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// Display prints all elements in the list
func (s *SingleList) Display() {
	if s.IsEmpty() {
		fmt.Println("empty list")
		return
	}

	current := s.Head
	fmt.Print("Head -> ")
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

// Delete removes the first node that matches the given value
func (s *SingleList) Delete(val int) {
	if s.IsEmpty() {
		fmt.Println("empty list")
		return
	}

	// Case: delete the head
	if s.Head.Value == val {
		s.Head = s.Head.Next
		return
	}

	prev := s.Head
	current := s.Head.Next

	for current != nil {
		if current.Value == val {
			prev.Next = current.Next
			return
		}
		prev = current
		current = current.Next
	}

	fmt.Printf("value %d not found in the list\n", val)
}

// Search checks whether a value exists in the list
func (s *SingleList) Search(val int) bool {
	current := s.Head
	for current != nil {
		if current.Value == val {
			return true
		}
		current = current.Next
	}
	return false
}

// Length returns the total number of nodes in the list
func (s *SingleList) Length() int {
	count := 0
	current := s.Head
	for current != nil {
		count++
		current = current.Next
	}
	return count
}

// func main() {
// 	list := NewSingleList()

// 	list.Append(10)
// 	list.Append(20)
// 	list.Append(30)
// 	list.Display() // Head -> 10 -> 20 -> 30 -> nil

// 	fmt.Println("Length:", list.Length())

// 	list.Delete(20)
// 	list.Display() // Head -> 10 -> 30 -> nil

// 	fmt.Println("Search 10:", list.Search(10))
// 	fmt.Println("Search 50:", list.Search(50))

// 	list.Delete(70) // not found
// 	list.Display()
// }
