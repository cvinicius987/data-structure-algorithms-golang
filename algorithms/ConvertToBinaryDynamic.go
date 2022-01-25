package algorithms

import (
	"log"
	"strconv"
	"strings"

	"github.com/cvinicius987/data-structure-algorithms-golang/stackdynamic"
)

func ConvertToBinaryDynamic(decimal int) string {

	stack := stackdynamic.StackDynamic{}

	for decimal > 0 {

		resto := decimal % 2

		stack.Push(resto)

		decimal = decimal / 2
	}

	var result strings.Builder

	for !stack.IsEmpty() {

		value, err := stack.Pop()

		if err != nil {
			log.Fatalln(err)
		}

		result.WriteString(strconv.Itoa(value))
	}

	return result.String()
}
