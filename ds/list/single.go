package list

import (
	"fmt"
	"strings"
)

// Singly Linked List

// single represents a singly linked list.
type single struct {
	head *ListNode
	size int
}

// NewSingle creates and returns an empty singly linked list.
func NewSingle() LinkedList {
	return &single{}
}

// IsEmpty reports whether the list has no elements.
func (s *single) IsEmpty() bool {
	return s.size == 0
}

// Len returns the number of elements in the list.
func (s *single) Len() int {
	return s.size
}

// Append adds a Value to the end of the list.
// Time Complexity: O(n)
func (s *single) Append(val int) {
	n := &ListNode{Value: val}

	if s.head == nil {
		s.head = n
		s.size = 1
		return
	}

	curr := s.head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = n
	s.size++
}

// Prepend adds a Value to the beginning of the list.
// Time Complexity: O(1)
func (s *single) Prepend(val int) {
	n := &ListNode{
		Value: val,
		Next:  s.head,
	}
	s.head = n
	s.size++
}

// Delete removes the first occurrence of val from the list.
// Returns true if deletion was successful.
func (s *single) Delete(val int) bool {
	if s.IsEmpty() {
		return false
	}

	// Case: delete head
	if s.head.Value == val {
		s.head = s.head.Next
		s.size--
		return true
	}

	prev := s.head
	curr := s.head.Next

	for curr != nil {
		if curr.Value == val {
			prev.Next = curr.Next
			s.size--
			return true
		}
		prev = curr
		curr = curr.Next
	}

	return false
}

// Search checks whether a Value exists in the list.
func (s *single) Search(val int) bool {
	for curr := s.head; curr != nil; curr = curr.Next {
		if curr.Value == val {
			return true
		}
	}
	return false
}

// Values returns all list elements as a slice snapshot.
func (s *single) Values() []int {
	Values := make([]int, 0, s.size)
	for curr := s.head; curr != nil; curr = curr.Next {
		Values = append(Values, curr.Value)
	}
	return Values
}

// String returns a readable representation of the list.
func (s *single) String() string {
	if s.IsEmpty() {
		return "empty list"
	}

	var b strings.Builder
	b.WriteString("Head -> ")

	for curr := s.head; curr != nil; curr = curr.Next {
		fmt.Fprintf(&b, "%d -> ", curr.Value)
	}

	b.WriteString("nil")
	return b.String()
}

func (s *single) Head() *ListNode {
	return s.head
}
