package linkedlist_circular

import "github.com/cvinicius987/data-structure-algorithms-golang/iterator"

type Node struct {
	element int
	previus *Node
	next    *Node
}

type CircularLinkedList struct {
	start *Node
}

func (l *CircularLinkedList) Add(element int) {

	node := Node{}

	node.element = element

	if l.IsEmpty() {

		node.previus = &node
		node.next = &node

		l.start = &node

	} else {

		tmp := l.start

		for tmp.next != l.start {
			tmp = tmp.next
		}

		l.start.previus = &node

		tmp.next = &node

		node.previus = tmp

		node.next = l.start
	}
}

func (l *CircularLinkedList) Iterator() iterator.Iterator {

	return &IteratorCircularLinkedList{
		coll: l,
	}
}

func (l *CircularLinkedList) IsEmpty() bool {

	return l.start == nil
}
