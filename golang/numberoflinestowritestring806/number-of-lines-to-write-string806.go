package main

import "fmt"

func main() {
	widths := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	S := "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(numberOfLines(widths[0:len(widths)], S))
}

func numberOfLines(widths []int, S string) []int {
	sum := 0
	count := 0

	for _, v := range S {
		val := widths[int(v)-97]
		if val+sum > 100 {
			count++
			sum = 0
		}
		sum += val
		// fmt.Println(v)
	}
	return []int{count + 1, sum}
}
