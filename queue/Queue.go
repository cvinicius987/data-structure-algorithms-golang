package queue

import "errors"

type Queue struct {
	elements []int
	current  int
}

func (q *Queue) Add(value int) {

	q.elements = append(q.elements, value)
	q.current++
}

func (q *Queue) Poll() (int, error) {

	if q.IsEmpty() {
		return -1, errors.New("Queue esta vazia")
	}

	q.current--

	return q.elements[q.current], nil
}

func (q *Queue) IsEmpty() bool {

	return q.current <= 0
}
