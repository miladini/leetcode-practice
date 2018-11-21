package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// fmt.Println(time.Microsecond)
	start := time.Now()
	// ... operation that takes 20 milliseconds ...

	// for c := 0; c < 1000000000; c++ {

	// }

	toLowerCase("ALKSJDLAKJSDHQHQUIWDHAKJSDHKJAHSDIUQWHDSx")
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)
}

func toLowerCase(str string) string {
	slc := make([]string, 0)
	for _, v := range str {
		// fmt.Println(v)
		// fmt.Println(reflect.TypeOf(v))
		slc = append(slc, strings.ToLower(string(v)))
	}
	fmt.Println(strings.Join(slc, ""))
	return strings.Join(slc, "")
}
