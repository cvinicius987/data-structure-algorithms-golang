package linkedlist

import (
	"testing"
)

func TestLinkedList_Poll(t *testing.T) {

	linkedList := LinkedList{}

	linkedList.Add(10)

	value, _ := linkedList.Poll()

	if value != 10 {
		t.Error("Value:", value, "Expected:", 10)
	}

	if !linkedList.IsEmpty() {
		t.Error("A LinkedList deveria estar vazia")
	}
}

func TestLinkedList_IsEmpty(t *testing.T) {

	linkedList := LinkedList{}

	linkedList.Add(1)

	linkedList.Poll()

	if !linkedList.IsEmpty() {
		t.Error("Queue deveria estar vazia")
	}
}
