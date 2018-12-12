package main

import "fmt"
import "strconv"

func main() {
	fmt.Println("HELLO")
	fmt.Println(convertToBase7(-700))
}

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	var sign string
	res := ""
	if num < 0 {
		num = num * -1
		sign = "-"
	} else {
		sign = ""
	}

	for num != 0 {
		res = strconv.Itoa(num%7) + res
		num = num / 7
	}

	return sign + res
}
