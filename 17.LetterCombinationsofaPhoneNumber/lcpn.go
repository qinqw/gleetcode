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
	dmap := make([]*map[string]byte, l)

	for i := 0; i < l; i++ {
		n, _ := strconv.Atoi(string(digits[i]))
		t := make(map[string]byte)
		for _, v := range lmap[n] {
			if i == 0 {
				t[string(v)] = 0
			} else {
				for k2 := range *dmap[i-1] {
					t[k2+string(v)] = 0
				}
			}
		}
		dmap[i] = &t
	}

	for k := range *dmap[l-1] {
		res = append(res, k)
	}
	return res
}
