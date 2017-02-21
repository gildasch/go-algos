package priorityqueue

type PriorityQueue struct {
	p *priorityq
}

func (pq PriorityQueue) Add(i, p int) {
	p = pq.p
	if p == nil {
		p = priorityq{i, p, 1, nil, nil}
		return
	}

	if p.priority < p {
		p.priority, p = p, p.priority
		p.head, i = i, p.head
	}

	p.size++
	if p.left.size < p.right.size {
		p.left.Add(i, p)
	} else {
		p.right.Add(i, p)
	}
}

func (pq PriorityQueue) Pop() (int, int) {
	// ... not so nice
}

type priorityq struct {
	head        int
	priority    int
	size        int
	left, right *PriorityQueue
}
