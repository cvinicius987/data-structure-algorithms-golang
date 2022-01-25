package search

import (
	"fmt"
)

func ExampleBinarySearch() {

	values := []int{}

	for i := 1; i <= 100; i++ {
		values = append(values, i)
	}

	//Resultado (Index, Iterações)
	fmt.Println(BinarySearch(values, 10))
	fmt.Println(BinarySearch(values, 150))
	//Output:
	//9 6
	//-1 7
}
