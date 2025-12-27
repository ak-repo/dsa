package list

import (
	"fmt"
	"strings"
)

// single -linedlist

type single struct {
	head *node
	size int
}

func NewList() LinkedList {
	return &single{}
}

func (s *single) IsEmpty() bool {
	return s.size == 0
}

func (s *single) Append(val int) {
	n := &node{value: val}
	if s.head == nil {
		s.head = n
		s.size = 1
		return
	}

	curr := s.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = n
	s.size++
}
func (s *single) Delete(val int) bool {
	if s.IsEmpty() {
		return false
	}

	if s.head.value == val {
		s.head = s.head.next
		s.size--
		return true
	}

	prev := s.head
	curr := s.head.next

	for curr != nil {

		if curr.value == val {
			prev.next = curr.next
			s.size--
			return true
		}
		prev = curr
		curr = curr.next
	}

	return false
}

func (s *single) Search(val int) bool {

	for curr := s.head; curr != nil; curr = curr.next {
		if curr.value == val {
			return true
		}

	}
	return false
}

func (s *single) Len() int {
	return s.size
}

func (s *single) Values() []int {
	values := make([]int, 0, s.size)
	for curr := s.head; curr != nil; curr = curr.next {
		values = append(values, curr.value)
	}
	return values

}

func (s *single) String() string {
	if s.IsEmpty() {
		return "empty list"
	}

	var b strings.Builder
	b.WriteString("Head -> ")

	for curr := s.head; curr != nil; curr = curr.next {
		fmt.Fprintf(&b, "%d -> ", curr.value)
	}

	b.WriteString("nil")
	return b.String()
}

// func (s *single) Append(val int) {}
