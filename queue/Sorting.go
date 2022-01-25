package queue

import (
	"log"
)

func Sorting(input []int) []int {

	queue := Queue{}

	temp := input

	for i := 0; i < len(input); i++ {

		currentIndexAdd := -1

		for j := 0; j < len(temp); j++ {

			if currentIndexAdd == -1 || temp[currentIndexAdd] < temp[j] {
				currentIndexAdd = j
			}
		}

		queue.Add(temp[currentIndexAdd])

		temp = append(temp[:currentIndexAdd], temp[currentIndexAdd+1:]...)
	}

	result := []int{}

	for !queue.IsEmpty() {

		value, err := queue.Poll()

		if err != nil {
			log.Fatalln(err)
		}

		result = append(result, value)
	}

	return result
}
