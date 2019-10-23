package lession1

import "fmt"

func CreateSli(x int) [][]int {
	s := make([][]int, x)
	for i := 0; i < x; i++ {
		s[i] = make([]int, x)
	}
	return s
}

func PullSli(s [][]int) [][]int {
	n := len(s)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				s[i][j] = 1
				s[i][n-1-i] = 1

			}
		}
	}
	return s
}

func PrintSli(s [][]int) {
	fmt.Println(s)
}
