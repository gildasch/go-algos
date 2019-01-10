package sort

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSortSimple(t *testing.T) {
	a := []int{4, 3, 6, 8, 4, 2, 5}

	QuickSort(a)

	assert.True(t, inOrder(a))
	assert.Len(t, a, 7)
}

func TestQuickSort(t *testing.T) {
	a := randomSlice(1000)

	QuickSort(a)

	assert.True(t, inOrder(a))
	assert.Len(t, a, 1000)
}

func randomSlice(n int) []int {
	var a []int

	for i := 0; i < n; i++ {
		a = append(a, rand.Intn(10000))
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
