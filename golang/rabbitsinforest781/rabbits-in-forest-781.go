package main

import (
	"fmt"
	"math"
)

func main() {
	answers := []int{10, 10, 10}
	fmt.Println(numRabbits(answers))
}

func numRabbits(answers []int) int {
	m := map[int]int{}
	for _, v := range answers {
		if _, ok := m[v]; ok {
			m[v] = m[v] + 1
		} else {
			m[v] = 1
		}
	}

	count := 0

	for k, v := range m {
		if k+1 == v {
			count += k + 1
		} else if k+1 > v {
			count += k + 1
		} else {
			count += (k + 1) * int(math.Ceil(float64(v)/float64(k+1)))
		}
	}

	return count
}
