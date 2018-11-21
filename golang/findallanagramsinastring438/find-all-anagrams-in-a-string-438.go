package main

import (
	"fmt"
)

func main() {
	fmt.Println(findAnagrams("abab", "ab"))
}

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	// res = append(res, 22)
	// res = append(res, 1)

	pMap := make(map[int]int)
	for i := 0; i < len(p); i++ {
		if v, ok := pMap[int(p[i])]; ok {
			pMap[int(p[i])] = v + 1
		} else {
			pMap[int(p[i])] = 1
		}
	}

	for i := 0; i+len(p)-1 < len(s); i++ {
		// res = append(res, int(s[i]))
		sMap := make(map[int]int)
		for j := 0; j < len(p); j++ {
			if v, ok := sMap[int(s[i+j])]; ok {
				sMap[int(s[i+j])] = v + 1
			} else {
				sMap[int(s[i+j])] = 1
			}
		}

		flag := true
		for k, v := range sMap {
			if pMap[k] != v {
				flag = false
			}
		}
		if flag {
			res = append(res, i)
		}
	}

	return res
}
