package sort

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortSimple(t *testing.T) {
	for name, f := range Algorithms {
		t.Run(name, func(t *testing.T) {
			a := []int{4, 3, 6, 8, 4, 2, 5}

			f(a)

			assert.True(t, inOrder(a))
			assert.Len(t, a, 7)
		})
	}
}

func TestSort(t *testing.T) {
	for name, f := range Algorithms {
		t.Run(name, func(t *testing.T) {
			a := randomSlice(1000)

			f(a)

			assert.True(t, inOrder(a))
			assert.Len(t, a, 1000)
		})
	}
}

func BenchmarkQuickSort(b *testing.B) {

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		a := randomSlice(1000)
		b.StartTimer()

		QuickSort(a)
	}
}

func BenchmarkMergeSort(b *testing.B) {

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		a := randomSlice(1000)
		b.StartTimer()

		MergeSort(a)
	}
}

func BenchmarkStdSort(b *testing.B) {

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		a := randomSlice(1000)
		b.StartTimer()

		sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	}
}

func BenchmarkQuickSort100000(b *testing.B) {

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		a := randomSlice(100000)
		b.StartTimer()

		QuickSort(a)
	}
}

func BenchmarkMergeSort100000(b *testing.B) {

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		a := randomSlice(100000)
		b.StartTimer()

		MergeSort(a)
	}
}

func BenchmarkStdSort100000(b *testing.B) {

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		a := randomSlice(100000)
		b.StartTimer()

		sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	}
}
