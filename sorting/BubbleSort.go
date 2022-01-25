package sorting

func BubbleSort(values *[]int) {

	logicBubbleSort(values, true)
}

func BubbleSortDesc(values *[]int) {

	logicBubbleSort(values, false)
}

func logicBubbleSort(values *[]int, asc bool) {

	fnDirection := func(intern int) bool {

		if asc {
			return (*values)[intern] > (*values)[intern+1]
		}

		return (*values)[intern] < (*values)[intern+1]
	}

	temp := 0

	for current := 0; current < len(*values)-1; current++ {

		navigation := (len(*values) - 1) - current

		for intern := 0; intern < navigation; intern++ {

			if fnDirection(intern) {

				//pega o valor corrente
				temp = (*values)[intern]

				//altera o valor corrente para o proximo
				(*values)[intern] = (*values)[intern+1]

				//Seta o proximo como sendo o corrente
				(*values)[intern+1] = temp
			}
		}
	}
}
