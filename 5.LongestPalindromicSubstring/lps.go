package main

import (
	"fmt"
)

func main() {
	s := "2skkwiksss"
	fmt.Println(longestPalindrome(s))

	s = "aaaaaa"
	fmt.Println(longestPalindrome(s))

}

func longestPalindrome(s string) string {
	fmt.Println(s)
	if len(s) == 0 {
		return ""
	}
	tmpMaxIndex := [2]int{}
	tmpIR := [128]int{}
	for i, c := range s {
		if tmpIR[c] > 0 {
			if i+1-tmpIR[c] > tmpMaxIndex[1]-tmpMaxIndex[0] {
				tmpMaxIndex[0] = tmpIR[c] - 1
				tmpMaxIndex[1] = i
			}
			//startIndex = i
		} else {
			tmpIR[c] = i + 1
		}
	}
	return s[tmpMaxIndex[0] : tmpMaxIndex[1]+1]
}
