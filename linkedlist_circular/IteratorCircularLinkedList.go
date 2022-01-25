package linkedlist_circular

import "errors"

type IteratorCircularLinkedList struct {
	coll    *CircularLinkedList
	current *Node
}

func (i *IteratorCircularLinkedList) Next() (interface{}, error) {

	if i.current == nil {
		return -1, errors.New("Iterator invalido")
	}

	return i.current.element, nil
}

func (i *IteratorCircularLinkedList) HasNext() bool {

	if i.current == nil {

		i.current = i.coll.start

		return true

	} else if i.current.next != i.coll.start {

		i.current = i.current.next

		return true
	}

	return false
}
