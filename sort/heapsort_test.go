package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapParentLeftRight(t *testing.T) {
	h := &heap{}

	assert.Equal(t, 0, h.parent(1))
	assert.Equal(t, 0, h.parent(2))
	assert.Equal(t, 1, h.parent(3))
	assert.Equal(t, 1, h.parent(4))
	assert.Equal(t, 2, h.parent(5))
	assert.Equal(t, 2, h.parent(6))

	assert.Equal(t, 1, h.left(0))
	assert.Equal(t, 2, h.right(0))
	assert.Equal(t, 3, h.left(1))
	assert.Equal(t, 4, h.right(1))
	assert.Equal(t, 5, h.left(2))
	assert.Equal(t, 6, h.right(2))
	assert.Equal(t, 9, h.left(4))
	assert.Equal(t, 10, h.right(4))
}

func TestMaxHeapify(t *testing.T) {
	h := &heap{
		a:   []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1},
		len: 10,
	}

	h.maxHeapify(1)
	assert.Equal(t, 10, h.len)
	assert.Equal(t, []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}, h.a)

	h = &heap{
		a:   []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1},
		len: 10,
	}
	h.maxHeapify(3)
	assert.Equal(t, 10, h.len)
	assert.Equal(t, []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}, h.a)

	h = &heap{
		a:   []int{8, 6, 3, 4, 2, 5, 4},
		len: 7,
	}
	h.maxHeapify(2)
	assert.Equal(t, 7, h.len)
	assert.Equal(t, []int{8, 6, 5, 4, 2, 3, 4}, h.a)
}
