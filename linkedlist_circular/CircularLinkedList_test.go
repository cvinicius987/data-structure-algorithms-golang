package linkedlist_circular

import (
	"fmt"
	"testing"
)

func TestLinkedList_Add(t *testing.T) {

	list := CircularLinkedList{}

	list.Add(10)
	list.Add(145)
	list.Add(74)

	iterator := list.Iterator()

	for i := 1; i <= 3; i++ {

		iterator.HasNext()

		value, _ := iterator.Next()

		if i == 1 && value != 10 {
			t.Error("Expected: ", 10, "Value: ", value)
		}

		if i == 2 && value != 145 {
			t.Error("Expected: ", 145, "Value: ", value)
		}

		if i == 3 && value != 74 {
			t.Error("Expected: ", 74, "Value: ", value)
		}
	}
}

func TestLinkedList_Iterator(t *testing.T) {

	list := CircularLinkedList{}

	list.Add(10)
	list.Add(11)
	list.Add(12)

	iterator := list.Iterator()

	for iterator.HasNext() {

		value, _ := iterator.Next()

		fmt.Println(value)
	}
}

func TestCircularLinkedList_IsEmpty(t *testing.T) {

	list := CircularLinkedList{}

	if !list.IsEmpty() {
		t.Error("A CircularLinkedList deveria estar vazia")
	}
}
