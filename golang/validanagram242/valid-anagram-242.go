package main

import "fmt"

func main() {
	fmt.Println(isAnagram("laamm", "mamal"))
}

func isAnagram(s string, t string) bool {
	map1 := make(map[string]int)
	map2 := make(map[string]int)
	for _, r := range s {
		map1[string(r)]++
	}

	for _, r := range t {
		map2[string(r)]++
	}

	for k := range map1 {
		if map1[k] != map2[k] {
			return false
		}
	}
	return true
}
