package main

import (
	"fmt"

	"github.com/ak-repo/dsa/ds/list"
	"github.com/ak-repo/dsa/problems"
)

func main() {

	// List problems
	li := list.NewSingle()
	li.Append(10)
	li.Append(20)
	li.Append(30)
	li.Append(40)
	li.Append(50)

	fmt.Println("linked list: ", li.String())

	//revers
	r := problems.ReverseLinkedlist(li.Head())
	fmt.Println("after reverse: ", r.Value)

}
