package main

import "fmt"

func check(str string) string {
	flag := false
	for i := 0; i < len(str)-1; i++ {
		if str[i] == str[i+1] {
			flag = true
			prev := str[:i]
			last := str[i+2:]
			str = prev + last
		}
	}
	if flag {
		return check(str)
	}
	return str
}
func Solution(S string) string {
	// write your code in Go 1.4
	return check(S)

}

func main() {
	fmt.Println(Solution("BABABA"))

}
