package main

import (
	"fmt"
)

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2}, {4, 3}}))
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))

	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}))
}

func spiralOrder(matrix [][]int) []int {
	var res []int
	m := len(matrix)
	if m == 0 {
		return res
	}
	n := len(matrix[0])
	res = make([]int, m*n)
	m1, m2 := -1, m
	n1, n2 := -1, n
	//1 right,2 down, 3 left ,4 up
	fn := 1
	i, j := 0, 0
	k := 0
	for n1 < j && j < n2 && m1 < i && i < m2 && k < m*n {
		//fmt.Printf("k:%d\n", k)
		res[k] = matrix[i][j]
		if n1 == n2-2 && m1 == m2-2 {
			break
		}
		switch fn {
		case 1:
			if j < n2-1 {
				j++
			} else {
				fn = 2
				if i < m2-1 {
					m1 = m1 + 1
					i++
				} else {
					break
				}
			}
		case 2:
			if i < m2-1 {
				i++
			} else {
				fn = 3
				if j > n1+1 {
					n2 = n2 - 1
					j--
				} else {
					break
				}
			}
		case 3:
			if j > n1+1 {
				j--
			} else {
				fn = 4
				if i > m1+1 {
					i--
					m2 = m2 - 1
				} else {
					break
				}
			}
		case 4:
			if i > m1+1 {
				i--
			} else {
				if j < n2-1 {
					fn = 1
					j++
					n1 = n1 + 1
				} else {
					break
				}
			}
		}
		k++
	}
	return res
}
