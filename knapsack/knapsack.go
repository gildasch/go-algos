package knapsack

import "fmt"

func Knapsack(capacity int, weights, profits []int) []int {
	total := map[string]int{}
	selection := map[string][]int{}

	for w := 0; w <= capacity; w++ {
		total[key(-1, w)] = 0
		selection[key(-1, w)] = nil
	}

	for i := 0; i < len(weights); i++ {
		for c := 0; c <= capacity; c++ {
			if weights[i] > c {
				total[key(i, c)], selection[key(i, c)] = total[key(i-1, c)], selection[key(i-1, c)]
				continue
			}

			t1, s1 := total[key(i-1, c-weights[i])]+profits[i], selection[key(i-1, c-weights[i])]
			t2, s2 := total[key(i-1, c)], selection[key(i-1, c)]

			if t1 >= t2 {
				total[key(i, c)], selection[key(i, c)] = t1, copyWith(s1, i)
				continue
			}

			total[key(i, c)], selection[key(i, c)] = t2, s2
		}
	}

	return selection[key(len(weights)-1, capacity)]
}

func key(i, c int) string { return fmt.Sprintf("%d_%d", i, c) }

func copyWith(a []int, i int) []int {
	ret := []int{}
	for _, v := range a {
		ret = append(ret, v)
	}
	ret = append(ret, i)
	return ret
}
