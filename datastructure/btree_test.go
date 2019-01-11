package datastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBTreeLen(t *testing.T) {
	assert.Equal(t, 3, BTree{5, 0, 0, 0, 1, 3, 0, 0}.len(0))
	assert.Equal(t, 2, BTree{5, 0, 0, 0, 1, 3, 0, 0}.len(4))
	assert.Equal(t, 0, BTree{0, 0, 0, 0}.len(0))
	assert.Equal(t, 1, BTree{5, 0, 0, 0}.len(0))
}

func TestBTreeGet(t *testing.T) {
	assert.Equal(t, 1, BTree{5, 0, 0, 0, 1, 3, 0, 0}.Get(0))
	assert.Equal(t, 3, BTree{5, 0, 0, 0, 1, 3, 0, 0}.Get(1))
	assert.Equal(t, 5, BTree{5, 0, 0, 0, 1, 3, 0, 0}.Get(2))
}
