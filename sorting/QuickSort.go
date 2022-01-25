package sorting

func QuickSort(values *[]int) {

	logicQuickSort(values, 0, len(*values)-1)
}

func logicQuickSort(values *[]int, start, end int) {

	var middle int

	if start > end {
		return
	}

	middle = partition(values, start, end)

	logicQuickSort(values, start, middle-1)
	logicQuickSort(values, middle+1, end)
}

func partition(values *[]int, init, end int) int {

	var ref, up, down, temp int

	ref = (*values)[init]
	down = init
	up = end

	for down < up {

		//encontra um numero maior que o pivo
		for (*values)[down] <= ref && down < end {
			down++ //avanço para encontrar um valor maior
		}

		for (*values)[up] > ref {
			up--
		}

		if down < up {
			temp = (*values)[down]

			(*values)[down] = (*values)[up]

			(*values)[up] = temp
		}
	}

	(*values)[init] = (*values)[up]

	(*values)[up] = ref

	return up
}
