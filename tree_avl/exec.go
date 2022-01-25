package tree_avl

import (
	"fmt"
	"strconv"
)

func Exec(arr []string) {

	//Rotação simples direita
	//10 5 15 12 20 25

	//Rotação dupla a direita
	//10 5 20 15 25 17

	//Rotação dupla a esquerda
	//10 5 20 15 25 17

	var tree *Tree

	for i := 1; i < len(arr); i++ {

		nextInt, _ := strconv.ParseUint(arr[i], 10, 32)

		if i == 1 {
			tree = &Tree{
				Element: &Element{int(nextInt)},
			}
		} else {
			tree = tree.Add(&Element{int(nextInt)})
		}

		tree.CalculeBalance()
		tree = tree.VerifyBalance()
	}

	fmt.Println(tree.PrintTree(0))
}
