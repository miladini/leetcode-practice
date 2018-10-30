package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Hello World")
	fmt.Println(findComplement(4))
}

func findComplement(num int) int {
	length := math.Floor(1 + math.Log2(float64(num)))
	return ((1 << uint(length)) - 1) ^ num
}
