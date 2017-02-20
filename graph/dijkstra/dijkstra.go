package dijkstra

import (
	"sort"
)

type Dijkstraer interface {
	Dist(a int) map[int]int
}

func ShortestPath(d Dijkstraer, beg, end int) int {
	return ShortestPathRec(d, beg, end, nil,
		&remaining{
			[]int{beg},
			map[int]int{beg: 0},
		})
}

func ShortestPathRec(d Dijkstraer, beg, end int, visited []int, remaining *remaining) int {
	ci, cd := remaining.pop()
	if ci == end {
		return cd
	} else if cd == -1 {
		return -1
	}

	hasBeenVisited := false
	for _, i := range visited {
		if i == ci {
			hasBeenVisited = true
			break
		}
	}

	if !hasBeenVisited {
		neighbours := d.Dist(ci)
		remaining.process(cd, neighbours)
		visited = append(visited, ci)
	}
	return ShortestPathRec(d, beg, end, visited, remaining)
}

type edge struct {
	id   int
	dist int
}

type remaining struct {
	order []int
	e     map[int]int
}

func (r *remaining) pop() (int, int) {
	if len(r.e) < 1 {
		return -1, -1
	}
	i := r.order[0]
	r.order = r.order[1:]
	defer delete(r.e, i)
	return i, r.e[i]
}

func (r *remaining) process(cd int, neighbours map[int]int) {
	for i, nd := range neighbours {
		nd += cd
		d, ok := r.e[i]
		if !ok {
			r.e[i] = nd
			r.order = append(r.order, i)
			continue
		}
		if nd < d {
			r.e[i] = nd
		}
	}

	sort.Sort(r)
}

func (r *remaining) Len() int {
	return len(r.e)
}

func (r *remaining) Swap(i, j int) {
	r.order[i], r.order[j] = r.order[j], r.order[i]
}

func (r *remaining) Less(i, j int) bool {
	return r.e[r.order[i]] < r.e[r.order[j]]
}
