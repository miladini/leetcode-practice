package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
func isUgly(num int) bool {

	if num <= 1 {
		return false
	} else if num == 1 {
		return true
	} else {
		for num != 1 && num%2 == 0 {
			num /= 2
		}
		for num != 1 && num%3 == 0 {
			num /= 3
		}
		for num != 1 && num%5 == 0 {
			num /= 5
		}
		return num == 1
	}
}
