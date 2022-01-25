package search

//SequenceSearch: Busca o registro de forma sequencial dentro da coleção
func SequenceSearch(values []int, valueToFind int) (int, int) {

	result := -1
	iterations := 0

	for i := 0; i < len(values); i++ {

		if valueToFind == values[i] {
			result = i
			iterations++
			break
		}

		iterations++
	}

	return result, iterations
}
