package list

type ListNode struct {
	Value int
	Prev  *ListNode
	Next  *ListNode
}

// Interface for all list types
type LinkedList interface {
	Append(val int)
	Prepend(val int)
	Delete(val int) bool
	Search(val int) bool
	Len() int
	IsEmpty() bool
	Values() []int
	String() string
	Head() *ListNode
}
