package sorting

func SelectionSort(values *[]int) {

	//variavel de troca
	var temp int

	//Iteração nos dados, excluindo um elemento do total
	for current := 0; current < len(*values)-1; current++ {

		//Item corrente da iteração
		initLoopValue := current

		//Proximo item da iteração
		nextLoopValue := current + 1

		//Começa a iteração a partir do proximo item da (initLoopValue)
		for i := initLoopValue + 1; i < len(*values); i++ {

			if (*values)[i] < (*values)[nextLoopValue] {
				nextLoopValue = i
			}
		}

		if (*values)[nextLoopValue] < (*values)[initLoopValue] {
			temp = (*values)[initLoopValue]

			(*values)[initLoopValue] = (*values)[nextLoopValue]

			(*values)[nextLoopValue] = temp
		}
	}
}
