package sorting

import "testing"

func TestSelectionSort_Asc(t *testing.T) {

	values := []int{50, 45, 14, 41}

	SelectionSort(&values)

	if values[0] != 14 {
		t.Error("Index:", 0, "Expected:", 14)
	}

	if values[1] != 41 {
		t.Error("Index:", 1, "Expected:", 41)
	}

	if values[2] != 45 {
		t.Error("Index:", 2, "Expected:", 45)
	}

	if values[3] != 50 {
		t.Error("Index:", 3, "Expected:", 50)
	}
}
