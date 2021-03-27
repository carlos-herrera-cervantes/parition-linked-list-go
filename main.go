package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	length int
	head   *node
	tail   *node
}

// Returns the length of linked list
func (list linkedList) getLength() int {
	return list.length
}

// Prints the elements of linked list
func (list linkedList) printList() {
	for list.head != nil {
		fmt.Printf("%v -> ", list.head.data)
		list.head = list.head.next
	}

	fmt.Println()
}

// Adds new node in the linked list
func (list *linkedList) pushBack(node *node) {
	if list.head == nil {
		list.head = node
		list.tail = node
		list.length++

		return
	}

	list.tail.next = node
	list.tail = node
	list.length++
}

// Split a linked list based on a value
func (list *linkedList) partitionByValue(value int) {
	node := list.head
	head := node
	tail := node

	for node != nil {
		next := node.next

		if node.data < value {
			node.next = head
			head = node
		} else {
			tail.next = node
			tail = node
		}

		node = next
	}

	tail.next = nil
	list.head = head
}

func main() {
	node1 := &node{data: 3}
	node2 := &node{data: 5}
	node3 := &node{data: 8}
	node4 := &node{data: 5}
	node5 := &node{data: 10}
	node6 := &node{data: 2}
	node7 := &node{data: 1}
	linkedList1 := linkedList{}

	linkedList1.pushBack(node1)
	linkedList1.pushBack(node2)
	linkedList1.pushBack(node3)
	linkedList1.pushBack(node4)
	linkedList1.pushBack(node5)
	linkedList1.pushBack(node6)
	linkedList1.pushBack(node7)

	linkedList1.printList()
	linkedList1.partitionByValue(5)

	linkedList1.printList()
}
