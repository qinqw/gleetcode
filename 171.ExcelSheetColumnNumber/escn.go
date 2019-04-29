package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(titleToNumber(""))
	fmt.Println(titleToNumber("A"))
	fmt.Println(titleToNumber("AA"))
	fmt.Println(titleToNumber("Z"))
	fmt.Println(titleToNumber("AZ"))
	fmt.Println(titleToNumber("ABZ"))
}

func titleToNumber(s string) int {
	res := 0
	// mapA := map[rune]int{
	// 	'A': 1, 'B': 2, 'C': 3, 'D': 4, 'E': 5,
	// 	'F': 6, 'G': 7, 'H': 8, 'I': 9, 'J': 10,
	// 	'K': 11, 'L': 12, 'M': 13, 'N': 14, 'O': 15,
	// 	'P': 16, 'Q': 17, 'R': 18, 'S': 19, 'T': 20,
	// 	'U': 21, 'V': 22, 'W': 23, 'X': 24, 'Y': 25,
	// 	'Z': 26,
	// }
	arrA := []int{
		'A': 1, 'B': 2, 'C': 3, 'D': 4, 'E': 5,
		'F': 6, 'G': 7, 'H': 8, 'I': 9, 'J': 10,
		'K': 11, 'L': 12, 'M': 13, 'N': 14, 'O': 15,
		'P': 16, 'Q': 17, 'R': 18, 'S': 19, 'T': 20,
		'U': 21, 'V': 22, 'W': 23, 'X': 24, 'Y': 25,
		'Z': 26,
	}

	for i := 0; i < len(s); i++ {
		//letter := rune(s[len(s)-1-i])
		res += int(math.Pow(float64(26), float64(i))) * arrA[s[len(s)-1-i]]
	}

	return res
}
