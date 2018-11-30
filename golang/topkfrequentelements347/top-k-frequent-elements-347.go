package main

import "fmt"

func main() {
	nums := [...]int{3, 3, 2, 0, 0, 2, 4, 4, 5, 5, 9, 9, 9}
	fmt.Println(topKFrequent(nums[:], 1))
	fmt.Println("HFAFASDA")
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	s := make([][]int, len(nums)+1)
	for i := range s {
		s[i] = make([]int, 0)
	}
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v] = m[v] + 1
		} else {
			m[v] = 1
		}
	}

	for k, v := range m {
		s[v] = append(s[v], k)
	}
	res := make([]int, 0)

	f := len(nums)

	for f > 0 && len(res) < k {
		if len(s[f]) > 0 {
			for _, v := range s[f] {
				res = append(res, v)
				if len(res) >= k {
					break
				}
			}
		}
		f--
	}
	return res
}
