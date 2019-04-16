package main

import (
	"fmt"
	"testing"
)

// TestLongestSubstringWithoutRepeatingCharacters is find the substring
func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	s := "abdiwhusssaa"
	fmt.Println(LongestSubstringWithoutRepeatingCharacters(s))
	s = "aaaaaa"
	fmt.Println(LongestSubstringWithoutRepeatingCharacters(s))

	s = "abcabcaa"
	fmt.Println(LongestSubstringWithoutRepeatingCharacters(s))

	s = " "
	fmt.Println(LongestSubstringWithoutRepeatingCharacters(s))

	s = ""
	fmt.Println(LongestSubstringWithoutRepeatingCharacters(s))
}

// LongestSubstringWithoutRepeatingCharacters2 is find the substring
func TestlengthOfLongestSubstring(t *testing.T) {
	s := "abdiwhusssaa"
	fmt.Println(lengthOfLongestSubstring(s))
	s = "aaaaaa"
	fmt.Println(lengthOfLongestSubstring(s))

	s = "abcabcaa"
	fmt.Println(lengthOfLongestSubstring(s))
	s = " "
	fmt.Println(lengthOfLongestSubstring(s))
}
