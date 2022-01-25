package hashing

import (
	"testing"
)

func TestMapHash(t *testing.T) {

	hash := MapHash{}

	hash.Put(1, "Kotlin")
	hash.Put(2, "Java")
	hash.Put(3, "Golang")

	result1, _ := hash.Get(1)

	if result1 != "Kotlin" {
		t.Error("Value:", result1, "Expected:", "Kotlin")
	}

	result2, _ := hash.Get(2)

	if result2 != "Java" {
		t.Error("Value:", result1, "Expected:", "Java")
	}

	result3, _ := hash.Get(3)

	if result3 != "Golang" {
		t.Error("Value:", result1, "Expected:", "Golang")
	}
}
