package main

import (
	"fmt"
)

func main() {

	fmt.Println(longestCommonPrefix([]string{"aca", "ac", "aba"}))

	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}

func longestCommonPrefix(strs []string) string {
	var res string
	count := len(strs)
	lMin := 0
	for i := 0; i < count; i++ {
		if i == 0 {
			lMin = len(strs[0])
		}
		if lMin > len(strs[i]) {
			lMin = len(strs[i])
		}
	}
	for i := 0; i < lMin && count > 0; i++ {
		t := 0
		for j := 0; j < count; j++ {
			t += int(strs[j][i])
			if (t % int(strs[j][i])) > 0 {
				return res
			}
		}
		res += string(strs[0][i])
	}
	return res
}
