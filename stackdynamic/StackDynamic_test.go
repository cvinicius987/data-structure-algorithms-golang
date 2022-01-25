package stackdynamic

import (
	"testing"
)

func TestStackDynamic_PushAndPop(t *testing.T) {

	stack := StackDynamic{}

	stack.Push(10)
	stack.Push(12)

	value, _ := stack.Pop()

	if value != 12 {
		t.Error("Value:", value, "Expected:", 12)
	}

	value, _ = stack.Pop()

	if value != 10 {
		t.Error("Value:", value, "Expected:", 10)
	}
}

func TestStackDynamic_IsEmpty(t *testing.T) {

	stack := StackDynamic{}

	stack.Push(10)
	stack.Push(12)

	stack.Pop()
	stack.Pop()

	if !stack.IsEmpty() {
		t.Error("StackDynamic deveria ser vazia")
	}
}
