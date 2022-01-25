package search

//BinarySearch: Os parametros devem ser uma coleção de dados ordenados
func BinarySearch(values []int, item int) (int, int) {

	start := 0
	middle := 0
	final := (len(values) - 1)

	iteractions := 0

	for start <= final {

		iteractions++

		middle = (start + final) / 2

		value := values[middle]

		if item == value {
			return middle, iteractions

		} else {

			if item > value {
				start = middle + 1
			} else {
				final = middle - 1
			}
		}
	}

	return -1, iteractions
}
