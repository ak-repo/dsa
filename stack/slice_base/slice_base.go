package main

import "fmt"

type Stack []int

// Push adds a value to the stack
func (s *Stack) Push(val int) {
	*s = append(*s, val)
}

// Pop removes and returns the top element
func (s *Stack) Pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top, true
}

// Peek returns the top element without removing it
func (s *Stack) Peek() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	return (*s)[len(*s)-1], true
}

// Len returns the stack size
func (s *Stack) Len() int {
	return len(*s)
}

func main() {
	var stack Stack

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println("Stack:", stack)

	if val, ok := stack.Peek(); ok {
		fmt.Println("Peek:", val)
	}

	if val, ok := stack.Pop(); ok {
		fmt.Println("Popped:", val)
	}

	fmt.Println("After pop:", stack)
	fmt.Println("Size:", stack.Len())
}
