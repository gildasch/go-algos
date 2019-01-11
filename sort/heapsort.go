package sort

func init() {
	Algorithms["heap sort"] = func(a []int) {
		Heapsort(a)
	}
}

func Heapsort(a []int) {
	h := buildHeap(a)

	for i := len(a) - 1; i >= 1; i-- {
		a[0], a[i] = a[i], a[0]
		h.len--
		h.maxHeapify(0)
	}
}

type heap struct {
	a   []int
	len int
}

func buildHeap(a []int) *heap {
	h := heap{a, len(a)}

	for i := len(a)/2 - 1; i >= 0; i-- {
		h.maxHeapify(i)
	}

	return &h
}

func (h *heap) maxHeapify(i int) {
	l, r := h.left(i), h.right(i)
	largest := i
	if l < h.len && h.a[l] > h.a[largest] {
		largest = l
	}
	if r < h.len && h.a[r] > h.a[largest] {
		largest = r
	}
	if largest != i {
		h.a[i], h.a[largest] = h.a[largest], h.a[i]
		h.maxHeapify(largest)
	}
}

func (h *heap) parent(i int) int { return (i+1)/2 - 1 }
func (h *heap) left(i int) int   { return 2*(i+1) - 1 }
func (h *heap) right(i int) int  { return 2 * (i + 1) }
