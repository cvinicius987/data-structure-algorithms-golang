package linkedlist

import (
	"errors"
)

type Node struct {
	element int
	next    *Node
}

type LinkedList struct {
	start *Node
}

func (l *LinkedList) Add(element int) {

	node := &Node{element, nil}

	if l.IsEmpty() {

		l.start = node

	} else {

		aux := l.start

		for aux.next != nil {
			aux = aux.next
		}

		aux.next = node
	}
}

func (l *LinkedList) Poll() (int, error) {

	if !l.IsEmpty() {

		result := l.start.element

		l.start = l.start.next

		return result, nil
	}

	return -1, errors.New("LinkedList esta vazia")
}

func (l *LinkedList) IsEmpty() bool {

	return l.start == nil
}
