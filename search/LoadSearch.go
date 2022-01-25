package search

import (
	"fmt"
)

func LoadSearch() {

	//Numero para buscar o index
	values := []int{}

	for i := 1; i <= 100; i++ {
		values = append(values, i)
	}

	//Numero para buscar o index
	numbers := []int{10, 25, 5, 89, 78, 150}

	for indexNum := 0; indexNum < len(numbers); indexNum++ {

		numberToFind := numbers[indexNum]

		fmt.Println(" \n ====== Numero: ", numberToFind)

		result, iteractions := SequenceSearch(values, numberToFind)

		fmt.Printf(" >>> SequenceSearch [Index: %v, Iterações: %v]\n", result, iteractions)

		result, iteractions = BinarySearch(values, numberToFind)

		fmt.Printf(" >>> SequenceBinary [Index: %v, Iterações: %v]\n", result, iteractions)
	}
}
