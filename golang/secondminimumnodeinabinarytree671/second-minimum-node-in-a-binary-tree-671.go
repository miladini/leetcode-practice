package main

import (
	"fmt"
	"math"
)

func main() {
	root := &TreeNode{Val: 2, Right: &TreeNode{Val: 2}, Left: &TreeNode{Val: 2}}

	fmt.Println(findSecondMinimumValue(root))

	fmt.Println("HELlO world!")
}

var smin = -1
var min = math.MaxInt32

func findSecondMinimumValue(root *TreeNode) int {
	smin = -1
	min = math.MaxInt32

	dfs(root)
	if smin == math.MaxInt32 {
		return -1
	}
	return smin
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Val < min {
		smin = min
		min = root.Val
	} else if root.Val < smin && root.Val != min {
		smin = root.Val
	}

	dfs(root.Left)
	dfs(root.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
