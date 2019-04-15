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
}

// LongestSubstringWithoutRepeatingCharacters2 is find the substring
func TestLongestSubstringWithoutRepeatingCharacters2(t *testing.T) {
	s := "abdiwhusssaa"
	fmt.Println(LongestSubstringWithoutRepeatingCharacters2(s))
	s = "aaaaaa"
	fmt.Println(LongestSubstringWithoutRepeatingCharacters2(s))

	s = "abcabcaa"
	fmt.Println(LongestSubstringWithoutRepeatingCharacters2(s))
}
