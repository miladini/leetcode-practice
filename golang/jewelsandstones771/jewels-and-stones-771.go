package main

import "fmt"

func main() {
	fmt.Println("salam!")
	fmt.Println(numJewelsInStones("salam", "lassssldj"))
}

func numJewelsInStones(J string, S string) int {
	jMap := make(map[rune]bool)
	for _, v := range J {
		jMap[v] = true
	}
	count := 0
	for _, v := range S {
		if jMap[v] {
			count++
		}
	}
	return count
}
