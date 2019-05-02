package main

import "fmt"

func main() {
	fmt.Println(reverseWords("the sky is blue"))
	fmt.Println(reverseWords("   the sky is blue"))
	fmt.Println(reverseWords("   the sky    is blue"))
	fmt.Println(reverseWords("   the sky    is blue  "))
	fmt.Println(reverseWords("blue"))
	fmt.Println(reverseWords("  blue"))
	fmt.Println(reverseWords("blue "))
}

func reverseWords(s string) string {
	res := ""
	j, k := len(s), len(s)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if j != k {
				if len(res) > 0 {
					res = res + " " + s[j:k]
				} else {
					res = s[j:k]
				}
			}
			j = i
			k = i
		} else {
			j = i
		}
	}

	if j != k {
		if len(res) > 0 {
			res = res + " " + s[j:k]
		} else {
			res = s[j:k]
		}

	}

	return string(res)
}
