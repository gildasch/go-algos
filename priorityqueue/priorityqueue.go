package priorityqueue

import (
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

type PriorityQueue[P constraints.Ordered, V any] struct {
	storage []node[P, V]
}

type node[P constraints.Ordered, V any] struct {
	value V
	prio  P
}

// p is the priority and must be a positive integer
func (pq *PriorityQueue[P, V]) Add(i V, p P) {
	pq.storage = append(pq.storage, node[P, V]{i, p})

	cur := len(pq.storage) - 1
	for cur > 0 {
		parent := cur / 2
		if pq.storage[parent].prio < pq.storage[cur].prio {
			pq.storage[parent], pq.storage[cur] = pq.storage[cur], pq.storage[parent]
		}
		cur = parent
	}
}

func (pq *PriorityQueue[P, V]) Pop() (value V, prio P) {
	if len(pq.storage) < 1 {
		return
	}

	ret := pq.storage[0]

	if len(pq.storage) == 1 {
		pq.storage = nil
		return ret.value, ret.prio
	}

	pq.storage[0] = pq.storage[len(pq.storage)-1]
	pq.storage = pq.storage[:len(pq.storage)-1]

	pq.MaxHeapify(0)

	return ret.value, ret.prio
}

func (pq *PriorityQueue[P, V]) MaxHeapify(cur int) {
	largest := cur

	left := 2*cur + 1
	right := 2*cur + 2

	if left < len(pq.storage) && pq.storage[largest].prio < pq.storage[left].prio {
		largest = left
	}

	if right < len(pq.storage) && pq.storage[largest].prio < pq.storage[right].prio {
		largest = right
	}

	if largest != cur {
		pq.storage[cur], pq.storage[largest] = pq.storage[largest], pq.storage[cur]
		pq.MaxHeapify(largest)
	}
}

func (pq *PriorityQueue[P, V]) String() string {
	content := []string{}
	for _, n := range pq.storage {
		content = append(content, fmt.Sprintf("<%v>", n))
	}
	return fmt.Sprintf("[ %s ]", strings.Join(content, " "))
}
