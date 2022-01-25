package hashing

type MapHash struct {
	values []interface{}
}

func (m MapHash) Put(key int, value interface{}) {

	hash := GenerateHash(key)

	m.values[hash] = value
}

func (m MapHash) Get(key int) (interface{}, error) {

	hash := GenerateHash(key)

	value := m.values[hash]

	return value, nil
}
