package sorting

import "fmt"

func InsertionSort(values *[]int) {

	logicInsertSort(values, true)
}

func InsertionSortDesc(values *[]int) {

	logicInsertSort(values, false)
}

func logicInsertSort(values *[]int, asc bool) {

	fnDirection := func(index int, valueFromVector int) bool {

		if asc {
			fmt.Println(">>> ", index, " >= 0 && ValorDoIndex [", index, "] > ", valueFromVector)

			return (index >= 0) && (*values)[index] > valueFromVector
		}

		fmt.Println(">>> ", index, " >= 0 && ValorDoIndex [", index, "] < ", valueFromVector)

		return (index >= 0) && (*values)[index] < valueFromVector
	}

	var indexLoopTemporario, valueFromVector int

	for indexLoopInicial := 1; indexLoopInicial < len(*values); indexLoopInicial++ {

		fmt.Println("\n> Array Atual:", *values)

		valueFromVector = (*values)[indexLoopInicial]

		fmt.Println(">> indexLoopInicial:", indexLoopInicial, "Value no Array:", valueFromVector)

		for indexLoopTemporario = indexLoopInicial - 1; fnDirection(indexLoopTemporario, valueFromVector); indexLoopTemporario-- {
			(*values)[indexLoopTemporario+1] = (*values)[indexLoopTemporario]

			fmt.Println(">>>> indexLoopTemporario:", indexLoopTemporario, ", Array Position [", indexLoopTemporario+1, "] = ", (*values)[indexLoopTemporario])
		}

		fmt.Println(">> Array Position[", indexLoopTemporario+1, "] = ", valueFromVector)

		(*values)[indexLoopTemporario+1] = valueFromVector
	}
}
