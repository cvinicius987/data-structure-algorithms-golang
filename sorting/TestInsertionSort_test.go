package sorting

import "testing"

func TestInsertionSort_Asc(t *testing.T) {

	values := []int{50, 45, 14, 41}

	InsertionSort(&values)

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

func TestInsertionSort_Desc(t *testing.T) {

	values := []int{50, 45, 14, 41}

	InsertionSortDesc(&values)

	if values[0] != 50 {
		t.Error("Index:", 0, "Expected:", 50)
	}

	if values[1] != 45 {
		t.Error("Index:", 1, "Expected:", 45)
	}

	if values[2] != 41 {
		t.Error("Index:", 2, "Expected:", 41)
	}

	if values[3] != 14 {
		t.Error("Index:", 3, "Expected:", 14)
	}
}
