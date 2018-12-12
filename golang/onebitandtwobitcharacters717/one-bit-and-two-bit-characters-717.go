package main

import "fmt"

func main() {
	fmt.Println("HELLO")
	inp := []int{1, 0, 0, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 1, 1, 1, 1, 0, 0}
	fmt.Println(isOneBitCharacter(inp[:]))
}

func isOneBitCharacter2(bits []int) bool {
	for i := 0; i < len(bits); i++ {
		if i == len(bits)-1 {
			return true
		} else if bits[i] == 1 {
			i++
		}
	}
	return false
}

func isOneBitCharacter(bits []int) bool {
	i := 0
	for ; i < len(bits)-1; i++ {
		if bits[i] == 1 {
			i++
		}
	}
	return i == len(bits)-1
}
