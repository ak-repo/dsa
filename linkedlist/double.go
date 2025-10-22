package main

import "fmt"

type DNode struct {
	Value int
	Next  *DNode
	Prev  *DNode
}

type DoubleList struct {
	Head *DNode
}

func NewDoubleList() *DoubleList {
	return &DoubleList{}
}

func (d *DoubleList) IsEmpty() bool {
	return d.Head == nil
}

func (d *DoubleList) InsertAtEnd(val int) {
	newNode := &DNode{Value: val}
	if d.IsEmpty() {
		d.Head = newNode
		return
	}

	current := d.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
	newNode.Prev = current
}

func (d *DoubleList) InsertAtBeginning(val int) {
	newNode := &DNode{Value: val}
	if d.IsEmpty() {
		d.Head = newNode
		return
	}

	d.Head.Prev = newNode
	newNode.Next = d.Head
	d.Head = newNode
}

func (d *DoubleList) Display() {
	if d.IsEmpty() {
		fmt.Println("empty list")
		return
	}

	current := d.Head
	fmt.Print("Head -> ")
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

func (d *DoubleList) DisplayBackward() {
	if d.IsEmpty() {
		fmt.Println("empty list")
		return
	}

	// Go to tail
	current := d.Head
	for current.Next != nil {
		current = current.Next
	}

	fmt.Print("Tail -> ")
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Prev
	}
	fmt.Println("nil")
}

func (d *DoubleList) Delete(val int) {
	if d.IsEmpty() {
		fmt.Println("empty list")
		return
	}

	current := d.Head
	for current != nil {
		if current.Value == val {
			if current.Prev != nil {
				current.Prev.Next = current.Next
			} else {
				d.Head = current.Next // head case
			}

			if current.Next != nil {
				current.Next.Prev = current.Prev
			}

			fmt.Println("value found and deleted")
			return
		}
		current = current.Next
	}
	fmt.Println("no value found")
}

func main() {
	dl := NewDoubleList()

	dl.InsertAtEnd(20)
	dl.InsertAtEnd(40)
	dl.InsertAtEnd(60)
	dl.InsertAtBeginning(10)
	dl.InsertAtBeginning(5)

	fmt.Println("Display forward:")
	dl.Display()

	fmt.Println("Display backward:")
	dl.DisplayBackward()

	dl.Delete(5)   // delete head
	dl.Delete(60)  // delete tail
	dl.Delete(100) // not found

	fmt.Println("After deletions:")
	dl.Display()
	dl.DisplayBackward()
}
