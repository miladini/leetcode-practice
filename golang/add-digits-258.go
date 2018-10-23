package main

import "fmt"

func main() {
	fmt.Println(addDigits(123456789))
}

func addDigits(num int) int {
	for num > 9 {
		sum := 0
		numc := num
		for numc > 0 {
			sum += numc % 10
			numc /= 10
		}

		num = sum
	}

	return num

}
