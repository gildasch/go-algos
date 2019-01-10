package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSortSimple(t *testing.T) {
	a := []int{4, 3, 6, 8, 4, 2, 5}

	MergeSort(a)

	assert.True(t, inOrder(a))
	assert.Len(t, a, 7)
}

func TestMergeSort(t *testing.T) {
	a := randomSlice(1000)

	MergeSort(a)

	assert.True(t, inOrder(a))
	assert.Len(t, a, 1000)
}
