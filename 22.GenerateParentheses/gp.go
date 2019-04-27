package main

import (
	"fmt"
)

func main() {
	//fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(4))
}

func generateParenthesis(n int) []string {
	pt := "()"
	res := []string{}
	mp := make([](map[string]int), n)
	for i := 0; i < n; i++ {
		if i == 0 {
			mp[0] = map[string]int{pt: 0}
			continue
		}
		tmpMp := make(map[string]int)
		for k := range mp[i-1] {
			for j := 0; j <= len(k); j++ {
				tmp := string(k[:j]) + pt + string(k[j:])
				tmpMp[tmp] = i
			}
			mp[i] = tmpMp
		}
	}
	//fmt.Println(mp)

	if n > 0 {
		for k := range mp[n-1] {
			res = append(res, k)
		}
	}

	return res
}
