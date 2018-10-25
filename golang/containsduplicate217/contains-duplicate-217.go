package main

import (
	"fmt"
)

func main() {
	inp := [...]int{2, 3, 1, 4, 4}
	fmt.Println(containsDuplicate(inp[0:len(inp)]))
}

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}
	return len(m) != len(nums)
}
