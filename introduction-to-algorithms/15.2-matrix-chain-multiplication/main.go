package main

import (
	"fmt"
	"time"
)

type Matrix struct {
	r, c int
}

func main() {
	matrices := []Matrix{
		{30, 35}, {35, 15}, {15, 5}, {5, 10}, {10, 20}, {20, 25},
	}

	min := minScalarMultiplications(matrices)

	start := time.Now()
	fmt.Printf("minimum of scalar multiplications: %d (in %v)\n", min, time.Since(start))
}

func minScalarMultiplications(matrices []Matrix) int {
	return minScalarMultiplicationsAux(matrices, map[string]int{})
}

func minScalarMultiplicationsAux(matrices []Matrix, cache map[string]int) int {
	request := fmt.Sprintf("%v", matrices)

	if min, ok := cache[request]; ok {
		return min
	}

	if len(matrices) == 0 || len(matrices) == 1 {
		return 0
	}

	if len(matrices) == 2 {
		return numberOfScalarMultiplications(matrices[0], matrices[1])
	}

	min := -1
	for splitAt := 1; splitAt < len(matrices); splitAt++ {
		minLeft := minScalarMultiplicationsAux(matrices[:splitAt], cache)
		minRight := minScalarMultiplicationsAux(matrices[splitAt:], cache)
		curr := minLeft + minRight + numberOfScalarMultiplications(
			Matrix{matrices[0].r, matrices[splitAt-1].c},
			Matrix{matrices[splitAt].r, matrices[len(matrices)-1].c},
		)

		if min == -1 || curr < min {
			min = curr
		}
	}

	cache[request] = min
	return min
}

func numberOfScalarMultiplications(m1, m2 Matrix) int {
	return m1.r * m1.c * m2.c
}
