package hashing

import (
	"fmt"
	"testing"
)

func TestGenerateHash(t *testing.T) {

	value := GenerateHash(1)

	fmt.Println(value)

}
