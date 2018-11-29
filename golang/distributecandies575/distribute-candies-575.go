package main

import "fmt"

func main() {
	candies := [...]int{1, 1, 2, 3}
	fmt.Println(distributeCandies(candies[:]))
}

func distributeCandies(candies []int) int {
	set := make(map[int]bool)
	for _, v := range candies {
		set[v] = true
	}
	return min(len(set), len(candies)/2)
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
