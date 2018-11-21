package main

import "fmt"

func main() {
	fmt.Println("salam")
}

func transpose(A [][]int) [][]int {
	res := make([][]int, 0)

	for j := 0; j < len(A[0]); j++ {
		col := make([]int, 0)
		for i := 0; i < len(A); i++ {
			col = append(col, A[i][j])
		}
		res = append(res, col)
	}
	return res
}
