package sorting

import "testing"

func TestQuickSort_Asc(t *testing.T) {

	values := []int{25, 57, 48, 37, 12, 92, 33}

	QuickSort(&values)

	if values[0] != 12 {
		t.Error("Index:", 0, "Expected:", 12)
	}

	if values[1] != 25 {
		t.Error("Index:", 1, "Expected:", 25)
	}

	if values[2] != 33 {
		t.Error("Index:", 2, "Expected:", 33)
	}

	if values[3] != 37 {
		t.Error("Index:", 3, "Expected:", 37)
	}

	if values[4] != 48 {
		t.Error("Index:", 4, "Expected:", 48)
	}

	if values[5] != 57 {
		t.Error("Index:", 5, "Expected:", 57)
	}

	if values[6] != 92 {
		t.Error("Index:", 6, "Expected:", 92)
	}
}
