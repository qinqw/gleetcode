package main

import (
	"fmt"
)

func main() {
	s := "abdiwhusssaa"
	fmt.Println(LongestSubstringWithoutRepeatingCharacters2(s))

	s = "aaaaaa"
	fmt.Println(LongestSubstringWithoutRepeatingCharacters2(s))

	s = "abcabcaa"
	fmt.Println(LongestSubstringWithoutRepeatingCharacters2(s))

	s = "pwwkew"
	fmt.Println(LongestSubstringWithoutRepeatingCharacters2(s))
}

// LongestSubstringWithoutRepeatingCharacters is find the substring
func LongestSubstringWithoutRepeatingCharacters(str string) int {
	res := 1
	maxLen := 0
	tmpMap := make(map[rune]byte)
	for _, c := range str {
		//fmt.Printf("[%d] %c\n", i, c)
		if _, ok := tmpMap[c]; ok {
			tmpMap := make(map[rune]byte)
			if res < maxLen {
				res = maxLen
			}
			maxLen = 1
			tmpMap[c] = 0
		} else {
			tmpMap[c] = 0
			maxLen++
		}

	}
	return res
}

// LongestSubstringWithoutRepeatingCharacters2 is find the substring
func LongestSubstringWithoutRepeatingCharacters2(str string) int {
	res := 1
	maxLen := 0
	tmpMap := make(map[byte]byte)
	for i := range str {
		//fmt.Printf("[%d] %c\n", i, str[i])
		if _, ok := tmpMap[str[i]]; ok {
			tmpMap := make(map[byte]byte)
			if res < maxLen {
				res = maxLen
			}
			maxLen = 1
			tmpMap[str[i]] = 0
		} else {
			tmpMap[str[i]] = 0
			maxLen++
		}

	}
	return res
}
