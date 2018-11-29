package main

import (
	"fmt"
)

func main() {
	fmt.Println("HELLOOO")
	fmt.Println(binaryGap(536870912))
	fmt.Println(binaryGap(536870911))
}

func binaryGap(N int) int {
	if 0 == (N & (N - 1)) {
		return 0
	} else {
		first := -1
		count := 0
		max := -1
		for N != 0 {
			if N&1 == 1 {
				if first == -1 {
					first = count
				} else {
					max = m(max, count-first)
					first = count
				}
			}
			N >>= 1
			count++
		}
		return max
	}
}

func m(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
