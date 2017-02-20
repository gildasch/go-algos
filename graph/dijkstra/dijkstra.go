package dijkstra

import "sort"

type Dijkstraer interface {
	Dist(a int) map[int]int
}

func ShortestPath(d Dijkstraer, beg, end int) int {
	return ShortestPathRec(d, beg, end, []*edge{&edge{beg, 0}})
}

func ShortestPathRec(d Dijkstraer, beg, end int, remaining []*edge) int {
	if len(remaining) < 1 {
		return -1
	}

	current := remaining[0]
	if current.id == end {
		return current.dist
	}

	remaining = remaining[1:]
	neighbours := d.Dist(current.id)
	for _, e := range remaining {
		if d, ok := neighbours[e.id]; ok && current.dist+d < e.dist {
			e.dist = current.dist + d
		}
	}
	sort.Sort(remainingSorter(remaining))

	return ShortestPathRec(d, beg, end, remaining)
}

type edge struct {
	id   int
	dist int
}

type remainingSorter []*edge

func (r remainingSorter) Len() int {
	return len(r)
}

func (r remainingSorter) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r remainingSorter) Less(i, j int) bool {
	return r[i].dist < r[j].dist
}
