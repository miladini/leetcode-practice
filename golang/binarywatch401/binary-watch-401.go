package main

import "fmt"

func main() {
	// fmt.Println("THIS Is IT")
	readBinaryWatch(2)
}
func readBinaryWatch(num int) []string {
	res := []string{}
	m := make(map[int][]int)

	for i := 0; i < 60; i++ {
		cb := countBits(i)
		m[cb] = append(m[cb], i)
	}
	// fmt.Println(m)

	for i := 0; i <= num; i++ {
		for _, hr := range m[i] {
			if hr > 11 {
				continue
			}
			for _, mn := range m[num-i] {
				resS := fmt.Sprintf("%d:%02d", hr, mn)
				res = append(res, resS)
				// fmt.Print(hr)
				// fmt.Print(":")
				// fmt.Print(mn)
				// fmt.Println()
				// fmt.Println(string(hr) + ":" + string(mn))
			}
		}
		// fmt.Println("----------")
	}
	fmt.Println(res)
	return res
}

func countBits(num int) int {
	res := 0
	for num != 0 {
		res = res + (num & 1)
		num >>= 1
	}
	return res
}
