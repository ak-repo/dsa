package main

import "fmt"

//
// StackNode represents a single node in the stack.
//
type StackNode struct {
	Value int
	Next  *StackNode
}

//
// Stack represents the stack data structure.
//
type Stack struct {
	Top  *StackNode
	Size int
}

//
// NewStack creates and returns an empty stack.
//
func NewStack() *Stack {
	return &Stack{}
}

//
// IsEmpty checks if the stack has no elements.
//
func (s *Stack) IsEmpty() bool {
	return s.Top == nil
}

//
// Push adds a new value to the top of the stack.
//
func (s *Stack) Push(val int) {
	newNode := &StackNode{Value: val, Next: s.Top}
	s.Top = newNode
	s.Size++
}

//
// Pop removes and returns the top element of the stack.
// Returns (0, false) if the stack is empty.
//
func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		fmt.Println("stack is empty")
		return 0, false
	}

	val := s.Top.Value
	s.Top = s.Top.Next
	s.Size--

	return val, true
}

//
// Peek returns the top element without removing it.
// Returns (0, false) if the stack is empty.
//
func (s *Stack) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	return s.Top.Value, true
}

//
// Display prints the stack elements from top to bottom.
//
func (s *Stack) Display() {
	if s.IsEmpty() {
		fmt.Println("stack is empty")
		return
	}

	fmt.Println("Stack (top → bottom):")
	current := s.Top
	for current != nil {
		fmt.Printf("-> %d\n", current.Value)
		current = current.Next
	}
	fmt.Println("Bottom")
}

//
// Len returns the number of elements in the stack.
//
func (s *Stack) Len() int {
	return s.Size
}

//
// main demonstrates stack operations.
//
// func main() {
// 	stack := NewStack()

// 	stack.Push(30)
// 	stack.Push(20)
// 	stack.Push(10)

// 	stack.Display()
// 	fmt.Printf("Current size: %d\n", stack.Len())

// 	if top, ok := stack.Peek(); ok {
// 		fmt.Println("Peek:", top)
// 	}

// 	if val, ok := stack.Pop(); ok {
// 		fmt.Println("Popped:", val)
// 	}
// 	if val, ok := stack.Pop(); ok {
// 		fmt.Println("Popped:", val)
// 	}

// 	stack.Display()
// 	fmt.Printf("Current size: %d\n", stack.Len())
// }
