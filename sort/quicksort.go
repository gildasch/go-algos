package sort

func init() {
	Algorithms["quick sort"] = func(a []int) {
		Quicksort(a)
	}
}

func Quicksort(a []int) {
	quicksort(a, 0, len(a)-1)
}

func quicksort(a []int, lo, hi int) {
	if lo < hi {
		p := partition(a, lo, hi)
		quicksort(a, lo, p-1)
		quicksort(a, p+1, hi)
	}
}

func partition(a []int, lo, hi int) int {
	pivot := a[hi]

	i := lo
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			if i != j {
				a[i], a[j] = a[j], a[i]
			}
			i++
		}
	}
	a[i], a[hi] = a[hi], a[i]

	return i
}
