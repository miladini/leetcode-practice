package main

import "fmt"

func main() {
	fmt.Println("Hello")
	inp := [...]int{2, 3, 5, 9, 2, 4, 3, 5, 9}

	singleNumber(inp[0:len(inp)])
}

func singleNumber(nums []int) int {
	xor := 0
	for _, v := range nums {
		//fmt.Println(v)
		xor = xor ^ v
	}
	fmt.Println(xor)
	return xor
}
