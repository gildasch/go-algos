package dijkstra

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type dGraph map[int][]edge

func (d dGraph) Dist(a int) map[int]int {
	ret := make(map[int]int)
	for _, e := range d[a] {
		ret[e.id] = e.dist
	}
	return ret
}

func TestShortestPath(t *testing.T) {
	g := dGraph{
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

	shortest := ShortestPath(g, 0, 9)
	fmt.Println("The shortest path is", shortest)
	assert.Equal(t, 487, shortest)
}
