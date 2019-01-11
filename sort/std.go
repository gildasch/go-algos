package sort

import "sort"

func init() {
	Algorithms["std sort.Slice"] = func(a []int) {
		sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	}
}
