package main

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func cal(colsum []int64, rowsum []int64, c, r int) int {
	cbef, caft, rbef, raft := int64(0), int64(0), int64(0), int64(0)
	for i := 0; i < c; i++ {
		cbef += colsum[i]
	}
	for i := c + 1; i < len(colsum); i++ {
		caft += colsum[i]
	}
	for i := 0; i < r; i++ {
		rbef += rowsum[i]
	}
	for i := r + 1; i < len(rowsum); i++ {
		raft += rowsum[i]
	}

	if cbef == caft && rbef == raft {
		return 1
	} else {
		return 0
	}
}
func Solution(A [][]int) int {
	// write your code in Go 1.4
	cols := make([]int64, len(A[0]))
	rows := make([]int64, len(A))

	var sum int64
	for j := 0; j < len(A[0]); j++ {
		sum = 0
		for i := 0; i < len(A); i++ {
			sum += int64(A[i][j])
		}
		cols[j] = sum
	}

	for i := 0; i < len(A); i++ {
		sum = 0
		for j := 0; j < len(A[0]); j++ {
			sum += int64(A[i][j])
		}
		rows[i] = sum
	}

	cnt := 0
	//cnt = cal(rows, cols, 1,1)
	for c := 0; c < len(A); c++ {
		for r := 0; r < len(A[0]); r++ {
			cnt += cal(rows, cols, c, r)
		}
	}
	return cnt
}
