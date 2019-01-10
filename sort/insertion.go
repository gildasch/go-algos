package sort

func init() {
	Algorithms["insertion sort"] = func(a []int) {
		Insertionsort(a)
	}
}

func Insertionsort(a []int) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0 && a[j-1] > a[j]; j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}
