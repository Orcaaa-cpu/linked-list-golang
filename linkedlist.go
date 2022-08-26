package main

import (
	"fmt"
	// "math/rand"
)

type Node struct {
	value int
	info  interface{}
	next  *Node
	key   interface{}
	prev  *Node
}

type List struct {
	head *Node
	tail *Node
}

func NewNode(value int, next *Node) *Node {
	var n Node
	n.value = value
	n.next = next
	return &n
}

func TraverseLinkedList(head *Node) {
	temp := head
	for temp != nil {
		fmt.Printf("%d ", temp.value)
		temp = temp.next
	}
	fmt.Println()
}

func (l *List) Insert(d interface{}) {
	list := &Node{info: d, next: nil}
	if l.head == nil {
		l.head = list
	} else {
		p := l.head
		for p.next != nil {
			p = p.next
		}
		p.next = list
	}
}

func (L *List) InsertV1(key interface{}) {
	list := &Node{
		next: L.head,
		key:  key,
	}
	if L.head != nil {
		L.head.prev = list
	}
	L.head = list

	l := L.head
	for l.next != nil {
		l = l.next
	}
	L.tail = l
}

func Show(l *List) {
	p := l.head
	for p != nil {
		fmt.Printf("-> %v", p.info)
		p = p.next
	}
}

func DeleteLastNode(head *Node) *Node {
	if head == nil {
		return head
	}
	temp := head
	for temp.next.next != nil {
		temp = temp.next
	}
	temp.next = nil
	return head
}

func DeleteAfterKthNode(head *Node, k int) *Node {
	// Delete after Kth node.
	if head == nil {
		return head
	}
	temp := head
	for temp != nil {
		if temp.value == k {
			temp.next = temp.next.next
		}
		temp = temp.next
	}
	return head
}

func (l *List) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

func Display(list *Node) {
	for list != nil {
		fmt.Printf("%v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

func (l *List) Reverse() {
	curr := l.head
	var prev *Node
	l.tail = l.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
	Display(l.head)
}

func main() {
	head := NewNode(30, NewNode(10, NewNode(40, NewNode(50, NewNode(60, nil)))))
	// before
	fmt.Printf("Input Linked list is: ")
	TraverseLinkedList(head)
	head = DeleteAfterKthNode(head, 50)
	fmt.Printf("Node %d is deleted.\nUpdated Linked List after deletion is:\n", 10)

	//after
	fmt.Printf("Input after Linked list is: ")
	TraverseLinkedList(head)

	// sl := List{}

	// for i := 0; i<6; i++{
	// sl.Insert(rand.Intn(100))
	// }
	// Show(&sl)

	// fmt.Println("")
	// fmt.Printf("--> %v ", *sl.head.next.next)
	// fmt.Println("")

	link := List{}
	link.InsertV1(5)
	link.InsertV1(9)
	link.InsertV1(13)
	link.InsertV1(22)
	link.InsertV1(28)
	link.InsertV1(36)

	// fmt.Println("\n==============================\n")
	fmt.Printf("Head: %v\n", link.head.key)
	fmt.Printf("Tail: %v\n", link.tail.key)
	link.Display()
	fmt.Println("")
	fmt.Printf("head: %v\n", link.head.key)
	fmt.Printf("tail: %v\n", link.tail.key)
	link.Reverse()
	// fmt.Println("\n==============================\n")

}
