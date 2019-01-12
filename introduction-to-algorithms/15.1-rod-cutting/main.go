package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	lengths = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	prices  = []int{1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
)

func main() {
	nStr := os.Args[1]
	n, err := strconv.Atoi(nStr)
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	rn := maxRevenue(n)
	memoDuration := time.Since(start)

	start = time.Now()
	rn2 := maxRevenueBottomUp(n)
	bottomUpDuration := time.Since(start)

	if rn != rn2 {
		fmt.Printf("memoization and bottom-up gave different results: %d and %d\n", rn, rn2)
		return
	}

	fmt.Printf("maximum revenue for a rod of length %d: %d (memoization: %v, bottom-up: %v)\n",
		n, rn, memoDuration, bottomUpDuration)
}

func maxRevenue(n int) int {
	return maxRevenueAux(n, map[int]int{})
}

func maxRevenueAux(n int, cache map[int]int) int {
	if max, ok := cache[n]; ok {
		return max
	}

	max := 0

	for i, l := range lengths {
		if l == n {
			max = prices[i]
		}
	}

	for i, l := range lengths {
		if l > n {
			continue
		}
		subMax := prices[i] + maxRevenue(n-l)
		if subMax > max {
			max = subMax
		}
	}

	cache[n] = max
	return max
}

func maxRevenueBottomUp(n int) int {
	bottomUpCache := map[int]int{}

	for k := 0; k <= n; k++ {
		max := 0

		for i, l := range lengths {
			if l > k {
				continue
			}
			subMax := prices[i] + bottomUpCache[k-l]
			if subMax > max {
				max = subMax
			}
		}

		bottomUpCache[k] = max
	}

	return bottomUpCache[n]
}
