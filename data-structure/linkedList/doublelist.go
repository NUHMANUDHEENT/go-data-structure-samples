package main

import "fmt"

type DoubleNode struct {
	Prev  *DoubleNode
	Next  *DoubleNode
	Value int
}
type DoubleLinkedList struct {
	Head *DoubleNode
	Tail *DoubleNode
}

func (dl *DoubleLinkedList) DoubleListInsert(value int) {
	newNode := &DoubleNode{Value: value, Next: nil, Prev: nil}
	if dl.Tail == nil {
		dl.Head = newNode
		dl.Tail = newNode
		return
	}
	dl.Tail.Next = newNode
	newNode.Prev = dl.Tail
	dl.Tail = newNode

}
func (dl *DoubleLinkedList) Delete(value int) {
	if dl.Head == nil {
		return
	}
	current := dl.Head
	for current != nil && current.Value != value{
			current = current.Next
	}
	if current.Next != nil {
		current.Next.Prev = current.Prev
		current.Prev.Next = current.Next
	} else {
		dl.Tail = current.Prev
	}
}
func main() {
	dll := &DoubleLinkedList{}
	dll.DoubleListInsert(1)
	dll.DoubleListInsert(2)
	dll.DoubleListInsert(3)
	dll.DoubleListInsert(4)
	dll.Delete(3)
	dll.Traversal()

}
func (dl *DoubleLinkedList) Traversal() {
	for current := dl.Head; current != nil; current = current.Next {
		fmt.Println(current.Value)
	}
}
