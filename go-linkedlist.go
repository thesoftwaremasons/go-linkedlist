package go_linkedlist

import "fmt"

type Node struct {
	data int
	next *Node
}
type DoubleNode struct {
	prev *DoubleNode
	data any
	next *DoubleNode
}

func NewDoubleNode(data any) *DoubleNode {
	return &DoubleNode{nil, data, nil}
}

func AddDoubleNode(list *DoubleNode, val int) *DoubleNode {
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

func AddNode(list *Node, val int) *Node {
	//1,5,7,10,49
	newNode := &Node{
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

func DeleteBeginingNode(list *Node) *Node {
	list = list.next
	return list

}
func DeleteFromEndOfNode(list *Node) *Node {
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
func DeleteFromPosition(list *Node, position int) *Node {
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

func ReverseArray(list *Node) *Node {
	prev := &Node{}
	middle := &Node{}
	next := list
	for next.next != nil {
		prev = middle
		middle = next
		next = middle.next
		middle.next = prev
	}
	return middle

}

func AddCircularLinkedList(list *Node, val int) *Node {
	//2.5,6,2,8

	head, temp := list, list
	//temp := list
	newNode := &Node{}
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
func AddInfrontCircularLinkedList(list *Node, val int) *Node {
	head := list

	temp := &Node{
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

func AddAtTheEndOfCircularList(list *Node, val int) *Node {
	head := list
	newNode := &Node{
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

func DeleteAtTheTopOfCircularList(list *Node) *Node {
	head := list
	for head != list.next {
		list = list.next
	}
	head = head.next
	list.next = head
	return head
}

func DeleteOfCircularListByPosition(list *Node, position int) *Node {

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

func DisplayCircularLinkedList(list *Node) {
	head := list
	for head != list.next {
		fmt.Println(list.data)
		list = list.next
	}
	fmt.Println(list.data)
}

func Display(list *Node) {
	temp := list
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}
func MergeLinkedList(firstNode *Node, secondNode *Node) *Node {
	holder, head := &Node{}, &Node{}
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
