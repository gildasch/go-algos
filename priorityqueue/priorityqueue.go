package priorityqueue

type PriorityQueue struct {
	Less func(int, int) bool
	p    priorityq
}

type priorityq []node

type node struct {
	id   interface{}
	prio int
}

// p is the priority and must be a positive integer
func (pq *PriorityQueue) Add(i interface{}, p int) {
	if pq.Less == nil {
		// Default Less function
		pq.Less = func(a, b int) bool {
			return a < b
		}
	}

	pq.p = append(pq.p, node{i, p})

	cur := len(pq.p) - 1
	for cur > 0 {
		parent := cur / 2
		if pq.Less(pq.p[parent].prio, pq.p[cur].prio) {
			pq.p[parent], pq.p[cur] = pq.p[cur], pq.p[parent]
		}
		cur = parent
	}
}

func (pq *PriorityQueue) Pop() (interface{}, int) {
	if len(pq.p) < 1 {
		return -1, -1
	}

	ret := pq.p[0]

	if len(pq.p) == 1 {
		pq.p = nil
		return ret.id, ret.prio
	}

	pq.p[0] = pq.p[len(pq.p)-1]
	pq.p = pq.p[:len(pq.p)-1]

	pq.MaxHeapify(0)

	return ret.id, ret.prio
}

func (pq *PriorityQueue) MaxHeapify(cur int) {
	largest := cur

	left := 2*cur + 1
	right := 2*cur + 2

	if left < len(pq.p) && pq.Less(pq.p[largest].prio, pq.p[left].prio) {
		largest = left
	}

	if right < len(pq.p) && pq.Less(pq.p[largest].prio, pq.p[right].prio) {
		largest = right
	}

	if largest != cur {
		pq.p[cur], pq.p[largest] = pq.p[largest], pq.p[cur]
		pq.MaxHeapify(largest)
	}
}
