package stack

import (
	"errors"
	"log"
)

type Stack struct {
	elements []int
	top      int
}

func (s *Stack) State() {
	log.Println("State:", s)
}

func (s *Stack) Pop() (int, error) {

	if s.top == 0 {
		return 0, errors.New("Stack não possui valores.")
	}

	s.top--

	return s.elements[s.top], nil
}

func (s *Stack) Push(value int) {

	s.elements = append(s.elements, value)
	s.top++
}

func (s *Stack) IsEmpty() bool {

	return s.top <= 0
}
