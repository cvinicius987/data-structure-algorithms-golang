package stack

import (
	"testing"
)

func TestStack_Pop(t *testing.T) {

	stack := Stack{}

	stack.Push(10)

	value, err := stack.Pop()

	if value != 10 {
		t.Error("Value:", value, "Expected:", 10)
	}

	if err != nil {
		t.Error(err)
	}
}

func TestStack_IsEmpty(t *testing.T) {

	stack := Stack{}

	stack.Push(1)

	stack.Pop()

	if !stack.IsEmpty() {
		t.Error("Stack deveria estarvazia")
	}
}
