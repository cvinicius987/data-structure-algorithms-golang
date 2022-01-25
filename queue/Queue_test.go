package queue

import (
	"testing"
)

func TestQueue_Poll(t *testing.T) {

	queue := Queue{}

	queue.Add(1)

	value, _ := queue.Poll()

	if value != 1 {
		t.Error("Value:", value, "Expected:", 1)
	}

	if !queue.IsEmpty() {
		t.Error("A Queue deveria estar vazia")
	}
}

func TestQueue_IsEmpty(t *testing.T) {

	queue := Queue{}

	queue.Add(1)

	queue.Poll()

	if !queue.IsEmpty() {
		t.Error("Queue deveria estar vazia")
	}
}
