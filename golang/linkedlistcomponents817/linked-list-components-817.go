package main

import "fmt"

func main() {

	head := &ListNode{Val: 44, Next: &ListNode{Val: 23, Next: &ListNode{Val: 22, Next: &ListNode{Val: 11, Next: nil}}}}

	G := [...]int{11, 22, 44, 23}

	numComponents(head, G[:])

}
func numComponents(head *ListNode, G []int) int {
	m := make(map[int]*ListNode)

	iNode := head
	for iNode != nil {
		m[iNode.Val] = iNode
		iNode = iNode.Next
	}

	count := 0

	mg := make(map[int]*ListNode, len(G))
	for _, v := range G {
		mg[v] = m[v]
	}

	for _, v := range G {
		next := mg[v].Next
		if next == nil {
			count++
		} else if _, ok := mg[next.Val]; !ok {
			count++
		}
	}

	fmt.Printf("Count is:\t%d\n", count)
	return count
}

type ListNode struct {
	Val  int
	Next *ListNode
}
