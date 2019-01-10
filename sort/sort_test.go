package sort

import (
	"fmt"
	"math/rand"
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

func BenchmarkSort(b *testing.B) {
	sizes := []int{1000, 100000}
	inputSlices := map[string]func(n int) []int{
		"random": randomSlice,
		"sorted": sortedSlice,
	}

	for _, size := range sizes {
		for sliceType, generator := range inputSlices {
			for name, f := range Algorithms {
				b.Run(fmt.Sprintf("%s on %s list of size %d", name, sliceType, size), func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						b.StopTimer()
						a := generator(size)
						b.StartTimer()

						f(a)
					}
				})
			}
		}
	}
}

func randomSlice(n int) []int {
	var a []int

	for i := 0; i < n; i++ {
		a = append(a, rand.Int())
	}

	return a
}

func sortedSlice(n int) []int {
	var a []int

	for i := 0; i < n; i++ {
		a = append(a, i)
	}

	return a
}

func inOrder(a []int) bool {
	if len(a) < 2 {
		return true
	}

	for i := 0; i <= len(a)-2; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}

	return true
}
