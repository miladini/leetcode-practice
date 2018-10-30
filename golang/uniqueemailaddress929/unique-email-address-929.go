package main

import (
	"fmt"
	"strings"
)

func main() {
	ss := "fafa" + "sss"
	fmt.Println(ss)
	inp := [...]string{"f.af.a@saa", "f..aa@asdasd", "f.af.aa@asdad", "fafa@asdas", "fafa@asda"}
	fmt.Println(numUniqueEmails(inp[0:len(inp)]))
}
func numUniqueEmails(emails []string) int {
	m := make(map[string]bool)
	for _, v := range emails {
		sEmails := strings.Split(v, "@")
		loc := strings.Split(sEmails[0], "+")[0]
		eLoc := make([]string, 0)

		for _, ch := range loc {
			if ch != '.' {
				eLoc = append(eLoc, string(ch))
			}
		}

		m[strings.Join(eLoc, "")+"@"+sEmails[1]] = true
	}
	return len(m)
}
