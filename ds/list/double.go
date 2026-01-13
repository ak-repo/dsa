package list

import (
	"fmt"
	"strings"
)

// doubly linked list (not exported)
type doublyList struct {
	head *ListNode
	tail *ListNode
	size int
}

// NewDoubly returns a new empty doubly linked list.
func NewDoubly() LinkedList {
	return &doublyList{}
}

// IsEmpty returns true if the list has no elements.
func (l *doublyList) IsEmpty() bool {
	return l.size == 0
}

// Len returns the number of elements in the list.
func (l *doublyList) Len() int {
	return l.size
}

// Append adds a Value to the end of the list.
func (l *doublyList) Append(val int) {
	n := &ListNode{Value: val}

	if l.head == nil {
		l.head = n
		l.tail = n
		l.size = 1
		return
	}

	n.Prev = l.tail
	l.tail.Next = n
	l.tail = n
	l.size++
}

// Prepend adds a Value to the beginning of the list.
func (l *doublyList) Prepend(val int) {
	n := &ListNode{Value: val}

	if l.head == nil {
		l.head = n
		l.tail = n
		l.size = 1
		return
	}

	n.Next = l.head
	l.head.Prev = n
	l.head = n
	l.size++
}

// Delete removes the first occurrence of val.
// Returns true if deletion was successful.
func (l *doublyList) Delete(val int) bool {
	for curr := l.head; curr != nil; curr = curr.Next {

		if curr.Value != val {
			continue
		}

		if curr.Prev != nil {
			curr.Prev.Next = curr.Next
		} else {
			l.head = curr.Next
		}

		if curr.Next != nil {
			curr.Next.Prev = curr.Prev
		} else {
			l.tail = curr.Prev
		}

		l.size--
		return true
	}
	return false
}

// Search checks if a Value exists in the list.
func (l *doublyList) Search(val int) bool {
	for curr := l.head; curr != nil; curr = curr.Next {
		if curr.Value == val {
			return true
		}
	}
	return false
}

// Values returns all elements as a slice (snapshot).
func (l *doublyList) Values() []int {
	Values := make([]int, 0, l.size)
	for curr := l.head; curr != nil; curr = curr.Next {
		Values = append(Values, curr.Value)
	}
	return Values
}

// String returns a readable representation of the list.
func (l *doublyList) String() string {
	if l.IsEmpty() {
		return "empty list"
	}

	var b strings.Builder
	b.WriteString("Head <-> ")

	for curr := l.head; curr != nil; curr = curr.Next {
		fmt.Fprintf(&b, "%d <-> ", curr.Value)
	}

	b.WriteString("nil")
	return b.String()
}

func (s *doublyList) Head() *ListNode {
	return s.head
}
