package sort

func MergeSort(s []int) []int {
	if len(s) < 2 {
		return s
	}

	a := MergeSort(s[:len(s)/2])
	b := MergeSort(s[len(s)/2:])

	for i := 0; i < len(s); i++ {
		if len(a) == 0 {
			s = append(s, b...)
			break
		}
		if len(b) == 0 {
			s = append(s, a...)
			break
		}

		if a[0] <= b[0] {
			s[i] = a[0]
			a = a[1:]
			continue
		}

		s[i] = b[0]
		b = b[1:]
	}

	return s
}
