package list

type node struct {
	value int
	next  *node
}

type dnode struct {
	value int
	prev  *dnode
	next  *dnode
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
}
