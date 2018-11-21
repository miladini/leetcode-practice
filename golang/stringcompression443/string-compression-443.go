package main

import "fmt"

func main() {
	fmt.Println("slaam!")
	a := 21
	b := fmt.Sprintf("%d", a)
	fmt.Println(b)

	chars := [...]byte{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'}
	sChars := string(chars[:])
	fmt.Println(sChars)
	res := compress(chars[:])
	ssChars := string(chars[:])
	fmt.Println(ssChars)
	fmt.Println(res)
}
func compress(chars []byte) int {
	reading, writing := 0, 0
	for reading < len(chars) {
		nextReading := reading + 1
		for nextReading < len(chars) && chars[nextReading] == chars[reading] {
			nextReading++
		}
		if nextReading-reading == 1 {
			chars[writing] = chars[reading]
			writing++
		} else {
			chars[writing] = chars[reading]
			writing++
			toWrite := fmt.Sprintf("%d", nextReading-reading)
			for nextWriting := writing; nextWriting-writing < len(toWrite); nextWriting++ {
				chars[nextWriting] = toWrite[nextWriting-writing]
			}
			writing = writing + len(toWrite)
		}
		reading = nextReading
	}
	return writing
}
