package main

import (
	"fmt"
)

func main() {
	fmt.Println("SALAM!!")

	input := [][]int{{1, 1, 1}, {2, 1, 2}, {3, 1, 3}}
	floodFill(input[:], 0, 0, 4)
	fmt.Println(input)

	// acc := int64(0)
	// // tt := time.Time{}
	// for z := 0; z < 100000; z++ {
	// 	LG := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 	L := LG.Int31() % 100
	// 	for L == 0 {
	// 		L = LG.Int31() % 100
	// 	}
	// 	input2 := make([][]int, L)
	// 	for i := range input2 {
	// 		input2[i] = make([]int, L)
	// 		for j := 0; j < int(L); j++ {
	// 			input2[i] = append(input2[i], int(LG.Int31()%10))
	// 		}
	// 	}
	// 	start := time.Now()
	// 	floodFill(input2[:], int(LG.Int31()%L), int(LG.Int31()%L), int(LG.Int31()%10))
	// 	// fmt.Println(time.Now().Sub(start))
	// 	acc += int64(time.Now().Sub(start))
	// 	// tt.Add(time.Now().Sub(start))
	// }
	// fmt.Println(time.Duration(acc))

}
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	visited := make([][]bool, len(image))
	for i := range visited {
		visited[i] = make([]bool, len(image[0]))
	}

	dfs2(&image, &visited, sr, sc, newColor, image[sr][sc])
	return image
}

func dfs(image *[][]int, visited *[][]bool, sr int, sc int, newColor int, sColor int) {
	(*visited)[sr][sc] = true
	(*image)[sr][sc] = newColor
	if sr-1 >= 0 && (*image)[sr-1][sc] == sColor && !(*visited)[sr-1][sc] {
		dfs(image, visited, sr-1, sc, newColor, sColor)
	}
	if sr+1 < len(*image) && (*image)[sr+1][sc] == sColor && !(*visited)[sr+1][sc] {
		dfs(image, visited, sr+1, sc, newColor, sColor)
	}
	if sc-1 >= 0 && (*image)[sr][sc-1] == sColor && !(*visited)[sr][sc-1] {
		dfs(image, visited, sr, sc-1, newColor, sColor)
	}
	if sc+1 < len((*image)[0]) && (*image)[sr][sc+1] == sColor && !(*visited)[sr][sc+1] {
		dfs(image, visited, sr, sc+1, newColor, sColor)
	}
}

func dfs2(image *[][]int, visited *[][]bool, sr int, sc int, newColor int, sColor int) {

	if sr < 0 || sc < 0 || sr >= len(*image) || sc >= len((*image)[0]) || (*image)[sr][sc] != sColor || (*visited)[sr][sc] {
		return
	}

	(*visited)[sr][sc] = true
	(*image)[sr][sc] = newColor

	dfs2(image, visited, sr-1, sc, newColor, sColor)
	dfs2(image, visited, sr+1, sc, newColor, sColor)
	dfs2(image, visited, sr, sc-1, newColor, sColor)
	dfs2(image, visited, sr, sc+1, newColor, sColor)
}
