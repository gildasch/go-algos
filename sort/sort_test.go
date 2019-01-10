package sort

import (
	"sort"
	"testing"
)

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
