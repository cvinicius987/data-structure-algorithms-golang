package stackdynamic

import "errors"

type Node struct {
	element int
	next    *Node
}

type StackDynamic struct {
	start *Node
}

func (s *StackDynamic) Pop() (int, error) {

	if s.start == nil {
		return -1, errors.New("StackDynamic não possui valores.")
	}

	element := s.start.element

	s.start = s.start.next

	return element, nil
}

func (s *StackDynamic) Push(value int) {

	current := s.start

	node := &Node{value, current}

	s.start = node
}

func (s *StackDynamic) IsEmpty() bool {

	return s.start == nil
}
