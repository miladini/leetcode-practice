package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("FAFA")
	fmt.Println(reverseOnlyLetters("!!!F-A*F!A?"))
}

func reverseOnlyLetters(S string) string {

	arr := strings.Split(S, "")

	length := len(arr)
	left, right := 0, len(S)-1

	for left < length && (!((arr[left][0] >= 'a' && arr[left][0] <= 'z') || (arr[left][0] >= 'A' && arr[left][0] <= 'Z'))) {
		left++
	}
	for right >= 0 && (!((arr[right][0] >= 'a' && arr[right][0] <= 'z') || (arr[right][0] >= 'A' && arr[right][0] <= 'Z'))) {
		right--
	}
	for left < right {
		tmp := arr[left]
		arr[left] = arr[right]
		arr[right] = tmp
		left++
		right--
		for left < length && (!((arr[left][0] >= 'a' && arr[left][0] <= 'z') || (arr[left][0] >= 'A' && arr[left][0] <= 'Z'))) {
			left++
		}
		for right >= 0 && (!((arr[right][0] >= 'a' && arr[right][0] <= 'z') || (arr[right][0] >= 'A' && arr[right][0] <= 'Z'))) {
			right--
		}
	}
	return strings.Join(arr, "")
}
