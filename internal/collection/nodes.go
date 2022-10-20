package collection

import (
	"fmt"
	"strconv"
)

type Node struct {
	next  *Node
	value int
}

type LinkedList struct {
	head  *Node
	tail  *Node
	count int
}

func NewNode(value int) (node *Node) {
	node = &Node{
		next:  nil,
		value: value,
	}
	return
}

func (linkedList *LinkedList) Add(value int) {
	if linkedList.count == 0 {
		linkedList.head = NewNode(value)
		linkedList.tail = linkedList.head
	} else {
		linkedList.tail.next = NewNode(value)
		linkedList.tail = linkedList.tail.next
	}
	linkedList.count++
}

func (linkedList *LinkedList) Contains(value int) bool {
	n := linkedList.head
	for n.next != nil {
		if n.value == value {
			return true
		}
		n = n.next
	}
	return false
}

func (linkedList *LinkedList) Get(index int) (bool, int) {
	if 0 <= index && index < linkedList.count {
		n := linkedList.head
		for i := 0; i < index; i++ {
			n = n.next
		}
		return true, n.value
	}
	return false, -1
}

func (linkedList *LinkedList) Reverse() {
	curr := linkedList.head
	var prev *Node
	var next *Node
	linkedList.tail = curr

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	linkedList.head = prev
}

func (linkedList *LinkedList) Print() {
	n := linkedList.head
	var str = ""
	for n != nil {
		str += strconv.Itoa(n.value)
		if n != linkedList.tail {
			str += "->"
		}
		n = n.next
	}
	fmt.Println(str + "\n")
}
