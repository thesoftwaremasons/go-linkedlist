package go_linkedlist

import "fmt"

type node struct {
	data int
	next *node
}
type DoubleNode struct {
	prev *DoubleNode
	data any
	next *DoubleNode
}

func NewDoubleNode(data any) *DoubleNode {
	return &DoubleNode{nil, data, nil}
}

func AddDoubleNode(list *DoubleNode, val any) *DoubleNode {
	//2,5,4
	head := list
	if list.next == nil {
		newNode := &DoubleNode{
			prev: head,
			data: val,
			next: nil,
		}
		list.next = newNode
	} else {
		for list.next != nil {
			list = list.next
		}
		newNode := &DoubleNode{
			prev: list,
			data: val,
			next: nil,
		}
		list.next = newNode
	}
	return head
}

func AddNode(list *node, val int) *node {
	//1,5,7,10,49
	newNode := &node{
		data: val,
		next: nil,
	}
	if list.next == nil {

		list.next = newNode
	} else {
		temp := list
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = newNode
	}
	return list
}
func DeleteCircularNode(list *DoubleNode, pos int) *DoubleNode {
	head := list
	count := 1
	if pos == 1 {
		head = list.next

	} else {
		for list.next != nil {

			if count == pos-1 {
				list.next = list.next.next
				break
			}
			list = list.next
			count++
		}
	}
	return head
}

func ReverseCicularNode(list *DoubleNode) *DoubleNode {
	var prev *DoubleNode = nil
	var middle *DoubleNode = nil
	next := list
	for next.next != nil {
		prev = middle
		middle = next
		next = next.next
		middle.next = prev
	}
	return middle
}
func DisplayDoubleLinkedList(list *DoubleNode) {
	for list.next != nil {
		fmt.Println(list.data)
		list = list.next
	}
	fmt.Println(list.data)
}

func DeleteBeginingNode(list *node) *node {
	list = list.next
	return list

}
func DeleteFromEndOfNode(list *node) *node {
	holder := list
	for list.next != nil {
		if list.next.next == nil {
			list.next = nil
			break
		}
		list = list.next
	}
	return holder
}
func DeleteFromPosition(list *node, position int) *node {
	holder := list
	count := 1
	for list.next != nil {
		if count == position-1 {
			list.next = list.next.next
			break
		}
		list = list.next
		count++
	}
	return holder
}

func ReverseArray(list *node) *node {
	prev := &node{}
	middle := &node{}
	next := list
	for next.next != nil {
		prev = middle
		middle = next
		next = middle.next
		middle.next = prev
	}
	return middle

}

func AddCircularLinkedList(list *node, val int) *node {
	//2.5,6,2,8

	head, temp := list, list
	//temp := list
	newNode := &node{}
	newNode.data = val

	if temp == temp.next {
		newNode.next = temp
		temp.next = newNode
	} else {
		for head != temp.next {
			temp = temp.next
		}

		temp.next = newNode
		newNode.next = head

	}

	return head
}
func AddInfrontCircularLinkedList(list *node, val int) *node {
	head := list

	temp := &node{
		data: val,
		next: head,
	}
	for head != list.next {
		list = list.next
	}
	list.next = temp
	list = temp

	return list
}

func AddAtTheEndOfCircularList(list *node, val int) *node {
	head := list
	newNode := &node{
		data: val,
		next: nil,
	}
	for head != list.next {
		list = list.next
	}
	newNode.next = list.next
	list.next = newNode
	return head
}

func DeleteAtTheTopOfCircularList(list *node) *node {
	head := list
	for head != list.next {
		list = list.next
	}
	head = head.next
	list.next = head
	return head
}

func DeleteOfCircularListByPosition(list *node, position int) *node {

	head := list
	count := 1
	for head != list.next {
		if count == position {
			list = list.next
			head.next = list
			break
		}
		list = list.next
		count++
	}
	if head == list.next {
		list = list.next
		list.next = head
	}
	return head
}

func DisplayCircularLinkedList(list *node) {
	head := list
	for head != list.next {
		fmt.Println(list.data)
		list = list.next
	}
	fmt.Println(list.data)
}

func Display(list *node) {
	temp := list
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}
func MergeLinkedList(firstNode *node, secondNode *node) *node {
	holder, head := &node{}, &node{}
	if firstNode.data < secondNode.data {
		holder = firstNode
		firstNode = firstNode.next
	} else {
		holder = secondNode
		secondNode = secondNode.next
	}
	head = holder
	for firstNode != nil && secondNode != nil {
		if firstNode.data < secondNode.data {
			holder.next = firstNode
			holder = firstNode
			firstNode = firstNode.next
		} else {
			holder.next = secondNode
			holder = secondNode
			secondNode = secondNode.next
		}
	}
	if firstNode != nil {
		holder.next = firstNode
	} else {
		holder.next = secondNode
	}
	return head
}
