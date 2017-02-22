package dijkstra

import (
	"github.com/gildasch/go-algos/priorityqueue"
)

// Dijkstraer is the interface that a graph must satisfy to be Dijkstraed
type Dijkstraer interface {
	// Neighbours returns a map of all the neighbours of node a
	// - the keys are the identifier of the nodes
	// - the values are the distance from a to that node
	Neighbours(a int) map[int]int
}

// ShortestPath returns the minimal distance from node `beg` to node `end` in graph `d`
func ShortestPath(d Dijkstraer, beg, end int) int {
	pq := &priorityqueue.PriorityQueue{
		Less: func(a, b int) bool {
			return a > b
		}}
	pq.Add(beg, 0)

	return shortestPathRec(d, beg, end, make(map[int]bool), pq)
}

func shortestPathRec(d Dijkstraer, beg, end int, visited map[int]bool, remaining *priorityqueue.PriorityQueue) int {
	// fmt.Println(remaining)
	ci, cd := remaining.Pop()
	if ci == end {
		return cd
	} else if cd == -1 {
		return -1
	}

	if !wasVisited(visited, ci) {
		visited[ci] = true
		neighbours := d.Neighbours(ci)
		for i, nd := range neighbours {
			remaining.Add(i, cd+nd)
		}
	}
	return shortestPathRec(d, beg, end, visited, remaining)
}

func wasVisited(visited map[int]bool, i int) bool {
	b, ok := visited[i]
	return ok && b
}
