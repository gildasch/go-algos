package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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

	rn := maxRevenue(n)
	fmt.Printf("maximum revenue for a rod of length %d: %d\n", n, rn)
}

func maxRevenue(n int) int {
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

	return max
}
