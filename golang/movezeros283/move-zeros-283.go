package main

import "fmt"

func main() {
	nums := [...]int{3, 4, 5, 0, 1, 11, 0, 9}
	q := nums[0:len(nums)]
	moveZeroes(q)
	fmt.Println("fafa")
	fmt.Println(q)

}

func moveZeroes(nums []int) {
	q := make([]int, 0)
	for i := 0; i < len(nums); {
		if nums[i] == 0 {
			q = append(q, i)
			i++
		} else if len(q) > 0 {
			zi := q[0]
			q = q[1:]
			tmp := nums[i]
			nums[i] = nums[zi]
			nums[zi] = tmp
		} else {
			i++
		}
	}
}
