package search

import "fmt"

func ExampleSequenceSearch() {

	values := []int{}

	for i := 1; i <= 100; i++ {
		values = append(values, i)
	}

	//Resultado (Index, Iterações)
	fmt.Println(SequenceSearch(values, 10))
	fmt.Println(SequenceSearch(values, 150))
	//Output:
	//9 10
	//-1 100
}
