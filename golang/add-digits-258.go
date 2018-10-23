package main

import "fmt"

func main() {
	fmt.Println(addDigits(12345))
}

func addDigits(num int) int {
	if num == 0 {
		return 0
	} else if (num % 9) > 0 {
		return num % 9
	} else {
		return 9
	}
}
