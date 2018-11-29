package main

import "fmt"

func main() {
	fmt.Println("HAFA")
	fn := [...]int{2, 3, 4, 5}
	f := [...]int{9, 12, 42, 2, 98, 5, 22, 23, 53, 3, 1}
	fmt.Println(nextGreaterElement(fn[:], f[:]))
}

func nextGreaterElement(findNums []int, nums []int) []int {
	m := make(map[int]int)
	stack := make([]int, 0)
	res := make([]int, 0)

	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] < nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			m[nums[i]] = -1
		} else {
			m[nums[i]] = stack[len(stack)-1]
		}

		stack = append(stack, nums[i])
	}

	for _, v := range findNums {
		res = append(res, m[v])
	}
	return res
}
