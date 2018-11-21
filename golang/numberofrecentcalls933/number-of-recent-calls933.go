package main

import "fmt"

func main() {
	fmt.Println("this")
	obj := Constructor()
	fmt.Println(obj.Ping(1))
	fmt.Println(obj.Ping(2))
	fmt.Println(obj.Ping(100))
	fmt.Println(obj.Ping(3000))
	fmt.Println(obj.Ping(3001))
	fmt.Println(obj.Ping(3002))
	fmt.Println(obj.Ping(3003))
	fmt.Println(obj.Ping(4003))
	fmt.Println(obj.Ping(9002))

}

//RecentCounter test asd
type RecentCounter struct {
	s1 []int
}

//Constructor fafa
func Constructor() RecentCounter {
	s1 := make([]int, 0)
	s1 = append(s1, 0)
	return RecentCounter{s1: s1}
}

//Ping fafa
func (e *RecentCounter) Ping(t int) int {
	e.s1 = append(e.s1, t)
	if t <= 3000 {
		return len(e.s1) - 1
	} else {
		lo := 0
		hi := len(e.s1) - 1
		for lo != hi {
			mid := (lo + hi + 1) / 2
			if e.s1[mid] > t-3000 {
				hi = mid - 1
			} else {
				lo = mid
			}
		}
		if e.s1[lo] == t-3000 {
			return len(e.s1) - lo

		} else {
			return len(e.s1) - lo - 1
		}
	}

}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
