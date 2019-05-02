package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{{1, 2}, {3, 4}}
	fmt.Println(matrix)
	rotate(matrix)
	fmt.Println(matrix)

	matrix = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(matrix)
	rotate(matrix)
	fmt.Println(matrix)

	//[[1,2,3,4,5],[6,7,8,9,10],[11,12,13,14,15],[16,17,18,19,20],[21,22,23,24,25]]

	//fmt.Println(rotate(matrix))
}

func rotate(matrix [][]int) {
	var l int
	l = len(matrix)
	pl := l - 1
	for j := 0; j < l/2; j++ {
		for i := 0; i < l/2; i++ {
			matrix[j][i], matrix[i][pl-j], matrix[pl-j][pl-i], matrix[pl-i][j] =
				matrix[pl-i][j], matrix[j][i], matrix[i][pl-j], matrix[pl-j][pl-i]
			//fmt.Printf("i,j:%d,%d\n", i, j)
		}
		if l%2 == 1 {
			matrix[j][l/2], matrix[l/2][l-1-j], matrix[l-1-j][l/2], matrix[l/2][j] =
				matrix[l/2][j], matrix[j][l/2], matrix[l/2][l-1-j], matrix[l-1-j][l/2]
		}
	}
}
