package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("2"))
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations("9"))
}

func letterCombinations(digits string) []string {
	var res []string
	if len(digits) < 1 {
		return res
	}

	lmap := []string{" ", "*", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	l := len(digits)
	var dmap *[]string
	for i := 0; i < l; i++ {
		n, _ := strconv.Atoi(string(digits[i]))
		t := []string{}
		for _, v := range lmap[n] {
			if i == 0 {
				t = append(t, string(v))
			} else {
				for _, k2 := range *dmap {
					t = append(t, k2+string(v))
				}
			}
		}
		dmap = &t
	}

	return *dmap
}
