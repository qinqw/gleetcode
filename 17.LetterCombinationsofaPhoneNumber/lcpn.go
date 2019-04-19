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

		var t []string
		ll := len(lmap[n])
		if i == 0 {
			t = make([]string, ll)
		} else {
			t = make([]string, ll*len(*dmap))
		}
		for j := 0; j < ll; j++ {
			if i == 0 {
				t[j] = string(lmap[n][j])
			} else {
				for i, k2 := range *dmap {
					t[i*ll+j] = k2 + string(lmap[n][j])
				}
			}
		}
		dmap = &t
	}

	return *dmap
}
