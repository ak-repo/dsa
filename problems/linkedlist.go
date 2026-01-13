package problems

import "github.com/ak-repo/dsa/ds/list"

// reverse a linedlist

func ReverseLinkedlist(head *list.ListNode) *list.ListNode {

	var prev *list.ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
