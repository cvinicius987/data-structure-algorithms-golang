package iterator

type Iterator interface {
	Next() (interface{}, error)

	HasNext() bool
}

type UserIterator interface {
	Iterator() Iterator
}
