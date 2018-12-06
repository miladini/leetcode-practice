package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("PRINT IT!")
	// g := [...]int{1, 2, 3, 2, 4, 2, 6, 3, 1, 6, 7, 2, 7, 2, 7, 2, 11}
	// s := [...]int{1, 1, 23, 4, 2, 6, 8, 3, 8, 12, 66, 3, 6, 3}
	g := [...]int{1, 2}
	s := [...]int{1, 2, 3}
	fmt.Println(findContentChildren(g[:], s[:]))
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	cookies, children, count := 0, 0, 0
	for cookies < len(s) && s[cookies] < g[children] {
		cookies++
	}

	for cookies < len(s) && children < len(g) {
		count++
		cookies++
		children++
		for children < len(g) && cookies < len(s) && s[cookies] < g[children] {
			cookies++
		}
	}

	return count
}
