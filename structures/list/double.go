package list

import (
	"fmt"
	"strings"
)

// doubly linked list (not exported)
type doublyList struct {
	head *dnode
	tail *dnode
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

// Append adds a value to the end of the list.
func (l *doublyList) Append(val int) {
	n := &dnode{value: val}

	if l.head == nil {
		l.head = n
		l.tail = n
		l.size = 1
		return
	}

	n.prev = l.tail
	l.tail.next = n
	l.tail = n
	l.size++
}

// Prepend adds a value to the beginning of the list.
func (l *doublyList) Prepend(val int) {
	n := &dnode{value: val}

	if l.head == nil {
		l.head = n
		l.tail = n
		l.size = 1
		return
	}

	n.next = l.head
	l.head.prev = n
	l.head = n
	l.size++
}

// Delete removes the first occurrence of val.
// Returns true if deletion was successful.
func (l *doublyList) Delete(val int) bool {
	for curr := l.head; curr != nil; curr = curr.next {

		if curr.value != val {
			continue
		}

		if curr.prev != nil {
			curr.prev.next = curr.next
		} else {
			l.head = curr.next
		}

		if curr.next != nil {
			curr.next.prev = curr.prev
		} else {
			l.tail = curr.prev
		}

		l.size--
		return true
	}
	return false
}

// Search checks if a value exists in the list.
func (l *doublyList) Search(val int) bool {
	for curr := l.head; curr != nil; curr = curr.next {
		if curr.value == val {
			return true
		}
	}
	return false
}

// Values returns all elements as a slice (snapshot).
func (l *doublyList) Values() []int {
	values := make([]int, 0, l.size)
	for curr := l.head; curr != nil; curr = curr.next {
		values = append(values, curr.value)
	}
	return values
}

// String returns a readable representation of the list.
func (l *doublyList) String() string {
	if l.IsEmpty() {
		return "empty list"
	}

	var b strings.Builder
	b.WriteString("Head <-> ")

	for curr := l.head; curr != nil; curr = curr.next {
		fmt.Fprintf(&b, "%d <-> ", curr.value)
	}

	b.WriteString("nil")
	return b.String()
}
