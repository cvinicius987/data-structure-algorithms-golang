package queue

import (
	"reflect"
	"testing"
)

func BenchmarkSorting(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Sorting([]int{14, 12, 36, 78, 254, 120, 23})
	}
}

func TestSorting(t *testing.T) {

	t.Run("Ordenação de 7 elementos", func(t *testing.T) {
		numbers := []int{14, 12, 36, 78, 254, 120, 23}

		got := Sorting(numbers)
		want := []int{12, 14, 23, 36, 78, 120, 254}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}
