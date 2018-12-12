package main

import (
	"fmt"
)

func main() {
	fmt.Println("HELLO")
	tokens := [...]int{100, 200, 300, 400}
	P := 200
	fmt.Println(bagOfTokensScore(tokens[:], P))
}

func bagOfTokensScore(tokens []int, P int) int {
	max := 0
	dfs(tokens, P, 0, &max, &map[int]bool{})
	return max
}

func dfs(tokens []int, power int, point int, max *int, seen *map[int]bool) {
	// fmt.Printf("%d %d %d\n", power, point, *max)
	// fmt.Println(*seen)
	if *max < point {
		*max = point
	}

	for i, v := range tokens {
		if ok := (*seen)[i]; !ok {
			if power >= v {
				(*seen)[i] = true
				dfs(tokens, power-v, point+1, max, seen)
				(*seen)[i] = false
			} else if point >= 1 {
				(*seen)[i] = true
				dfs(tokens, power+v, point-1, max, seen)
				(*seen)[i] = false
			}
		}
	}
}
