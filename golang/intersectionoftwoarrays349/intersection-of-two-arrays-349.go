package main

import "fmt"

func main() {
	inp1 := [...]int{2, 3, 1, 4, 4}
	inp2 := [...]int{2, 3, 1, 4, 4}
	fmt.Println("Hello World")
	intersection(inp1[0:len(inp1)], inp2[0:len(inp2)])
}

func intersection(nums1 []int, nums2 []int) []int {
	map1 := make(map[int]bool)
	map2 := make(map[int]bool)
	for _, v := range nums1 {
		map1[v] = true
	}

	for _, v := range nums1 {
		map2[v] = true
	}

	ret := make([]int, 0)
	for v := range map1 {
		if _, ok := map2[v]; ok {
			append(ret, v)
		}
	}

	return ret
}
