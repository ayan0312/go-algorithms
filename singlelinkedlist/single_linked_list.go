package singlelinkedlist

import (
	"errors"
)

type Node struct {
	data interface{}
	next *Node
}

type SingleLinkedList struct {
	head *Node
	last *Node
	size int
}

func CreateNode(data interface{}) *Node {
	n := new(Node)
	n.data = data
	n.next = nil
	return n
}

func New() *SingleLinkedList {
	sll := new(SingleLinkedList)
	sll.size = 0
	return sll
}

func (sll *SingleLinkedList) Insert(data interface{}, index int) error {
	err := sll.CheckSize(index)
	if err != nil {
		return err
	}

	newNode := CreateNode(data)
	if sll.size == 0 {
		sll.head = newNode
		sll.last = newNode
	} else if index == 0 {
		newNode.next = sll.head
		sll.head = newNode
	} else if sll.size == index {
		sll.last.next = newNode
		sll.last = newNode
	} else {
		prevNode, _ := sll.Node(index - 1)
		newNode.next = prevNode.next
		prevNode.next = newNode.next
	}
	sll.size++
	return nil
}

func (sll *SingleLinkedList) Remove(index int) (*Node, error) {
	err := sll.CheckSize(index)
	if err != nil {
		return nil, err
	}
	var removedNode *Node
	if index == 0 {
		removedNode = sll.head
		sll.head = sll.head.next
	} else if index == sll.size-1 {
		prevNode, _ := sll.Node(index - 1)
		nextNode := prevNode.next.next
		removedNode = prevNode.next
		prevNode.next = nextNode
	}
	sll.size--
	return removedNode, nil
}

func (sll *SingleLinkedList) Node(index int) (*Node, error) {
	err := sll.CheckSize(index)
	if err != nil {
		return nil, err
	}

	temp := sll.head
	for i := 0; i < index; i++ {
		temp = temp.next
	}

	return temp, nil
}

func (sll *SingleLinkedList) CheckSize(index int) error {
	if index < 0 || index > sll.size {
		return errors.New("index out of size")
	}
	return nil
}
