package main

import (
	"fmt"
	"sort"
)

func Solution3(A []int) int {
	// write your code in Go 1.4
	if len(A) < 2 {
		return 0
	}
	sort.Ints(A)
	sum := 0
	merge := A[0] + A[1]
	for i := 2; i < len(A); i++ {
		sum += merge
		merge = A[i] + merge
	}
	sum += merge
	return sum
}

func main() {
	arr := []int{1000, 250, 100}
	fmt.Println(Solution3(arr))

}
