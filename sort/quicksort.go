package sort

func QuickSort(a []int) {
	quickSort(a, 0, len(a)-1)
}

func quickSort(a []int, lo, hi int) {
	if lo < hi {
		p := partition(a, lo, hi)
		quickSort(a, lo, p-1)
		quickSort(a, p+1, hi)
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
