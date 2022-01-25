package algorithms

import "fmt"

func ExampleConvertToBinary() {

	fmt.Println(ConvertToBinary(10))
	fmt.Println(ConvertToBinary(145))
	fmt.Println(ConvertToBinary(1425))
	fmt.Println(ConvertToBinary(45254))
	//Output:
	//1010
	//10010001
	//10110010001
	//1011000011000110
}
