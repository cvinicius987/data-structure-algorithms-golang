package algorithms

import "fmt"

func ExampleConvertToBinaryDynamic() {

	fmt.Println(ConvertToBinaryDynamic(10))
	fmt.Println(ConvertToBinaryDynamic(145))
	fmt.Println(ConvertToBinaryDynamic(1425))
	fmt.Println(ConvertToBinaryDynamic(45254))
	//Output:
	//1010
	//10010001
	//10110010001
	//1011000011000110
}
