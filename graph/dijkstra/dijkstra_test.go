package dijkstra

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type dGraph map[int][]edge

type edge struct {
	id   int
	dist int
}

func (d dGraph) Neighbours(a int) map[int]int {
	ret := make(map[int]int)
	for _, e := range d[a] {
		ret[e.id] = e.dist
	}
	return ret
}

var gSmall = dGraph{
	0: []edge{
		edge{1, 7},
		edge{2, 9},
		edge{3, 14},
	},
	1: []edge{
		edge{0, 7},
		edge{2, 10},
		edge{4, 15},
	},
	2: []edge{
		edge{0, 9},
		edge{1, 10},
		edge{3, 2},
		edge{4, 11},
	},
	3: []edge{
		edge{0, 14},
		edge{2, 2},
		edge{5, 9},
	},
	4: []edge{
		edge{1, 15},
		edge{2, 11},
		edge{5, 6},
	},
	5: []edge{
		edge{3, 9},
		edge{4, 6},
	},
}

var g = dGraph{
	0: []edge{
		edge{1, 85},
		edge{2, 217},
		edge{4, 173},
	},
	1: []edge{
		edge{0, 85},
		edge{5, 80},
	},
	2: []edge{
		edge{0, 217},
		edge{6, 186},
		edge{7, 103},
	},
	3: []edge{
		edge{7, 183},
	},
	4: []edge{
		edge{0, 173},
		edge{9, 502},
	},
	5: []edge{
		edge{1, 80},
		edge{8, 250},
	},
	6: []edge{
		edge{2, 186},
	},
	7: []edge{
		edge{2, 103},
		edge{3, 183},
		edge{9, 167},
	},
	8: []edge{
		edge{5, 250},
		edge{9, 84},
	},
	9: []edge{
		edge{4, 502},
		edge{7, 167},
		edge{8, 84},
	},
}

var gCorrupt = dGraph{
	0: []edge{
		edge{1, 7},
		edge{2, 9},
		edge{3, 14},
	},
	1: []edge{
		edge{0, 7},
		edge{4, 20},
	},
	2:  []edge{edge{10, 1}},
	4:  []edge{},
	10: []edge{edge{4, 100}},
}

func TestShortestPath(t *testing.T) {
	shortestSmall := ShortestPath(gSmall, 0, 5)
	assert.Equal(t, 20, shortestSmall)

	shortest := ShortestPath(g, 0, 9)
	assert.Equal(t, 487, shortest)
}

func TestNoPath(t *testing.T) {
	ret := ShortestPath(gSmall, 0, 6)
	assert.Equal(t, -1, ret)

	ret = ShortestPath(g, 0, 10)
	assert.Equal(t, -1, ret)
}

func TestBadBegOrEnd(t *testing.T) {
	ret := ShortestPath(gSmall, -1, 5)
	assert.Equal(t, -1, ret)

	ret = ShortestPath(g, 0, -1)
	assert.Equal(t, -1, ret)
}

func TestBadGraph(t *testing.T) {
	ret := ShortestPath(gCorrupt, 0, 4)
	assert.Equal(t, 27, ret)
}

func BenchmarkShortestPathSmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ShortestPath(gSmall, 0, 5)
	}
}

func BenchmarkShortestPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ShortestPath(g, 0, 9)
	}
}
