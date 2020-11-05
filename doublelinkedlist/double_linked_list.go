package doublelinkedlist

import "errors"

type Node struct {
	Data interface{}
	Next *Node
	Prev *Node
}

type DoubleLinkedList struct {
	Head *Node
	Tail *Node
	size int
}

func CreateNode(data interface{}) *Node {
	n := new(Node)
	n.Data = data
	return n
}

func New() *DoubleLinkedList {
	ll := new(DoubleLinkedList)
	ll.size = 0
	return ll
}

func (ll *DoubleLinkedList) Count() int {
	return ll.size
}

func (ll *DoubleLinkedList) IsEmpty() bool {
	return ll.size == 0
}

func (ll *DoubleLinkedList) Insert(data interface{}, index int) (*Node, error) {
	err := ll.CheckSize(index)
	if err != nil {
		return nil, err
	}

	var newNode *Node

	if ll.size == 0 || index == 0 {
		newNode = ll.Prepend(data)
	} else if ll.size-1 == index {
		newNode = ll.Append(data)
	} else {
		newNode = CreateNode(data)
		next, _ := ll.Node(index)
		prev := next.Prev

		prev.Next = newNode
		newNode.Prev = prev

		next.Prev = newNode
		newNode.Next = next

		ll.size++
	}

	return newNode, nil
}

func (ll *DoubleLinkedList) Delete(index int) (*Node, error) {
	err := ll.CheckSize(index)
	if err != nil {
		return nil, err
	}

	var deleted *Node

	if index == 0 {
		deleted = ll.Head
		ll.Head = ll.Head.Next
	} else {
		node, _ := ll.Node(index)
		prev := node.Prev
		next := node.Next

		prev.Next = next
		next.Prev = prev

		deleted = node
	}

	ll.size--
	return deleted, nil
}

func (ll *DoubleLinkedList) Append(data interface{}) *Node {
	newNode := CreateNode(data)

	if ll.size == 0 {
		ll.Head = newNode
		ll.Tail = ll.Head
	} else {
		ll.Tail.Next = newNode
		newNode.Prev = ll.Tail
		ll.Tail = newNode
	}

	ll.size++
	return newNode
}

func (ll *DoubleLinkedList) Prepend(data interface{}) *Node {
	newNode := CreateNode(data)

	if ll.size == 0 {
		ll.Head = newNode
		ll.Tail = ll.Head
	} else {
		ll.Head.Prev = newNode
		newNode.Next = ll.Head
		ll.Head = newNode
	}

	ll.size++
	return newNode
}

func (ll *DoubleLinkedList) Node(index int) (*Node, error) {
	err := ll.CheckSize(index)
	if err != nil {
		return nil, err
	}

	var temp *Node

	if ll.size-index < index {
		temp = ll.Tail
		for i := ll.size; i >= index; i-- {
			temp = temp.Prev
		}
	} else {
		temp = ll.Head
		for i := 0; i < index; i++ {
			temp = temp.Next
		}
	}

	return temp, nil
}

func (ll *DoubleLinkedList) CheckSize(index int) error {
	if index < 0 || index > ll.size {
		return errors.New("index out of size")
	}
	return nil
}
