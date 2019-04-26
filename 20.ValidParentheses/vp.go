package main

import (
	"fmt"
)

func main() {
	s := "()"
	s = string(s[:0] + s[2:])
	fmt.Printf(s)
	fmt.Println(isValid("()"))

	fmt.Println(isValid("[()]"))
	fmt.Println(isValid("[([)]]"))

	fmt.Println(isValid("[({{{{{{{{{{}}}}}}}}}})]"))
}

func isValid(s string) bool {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		v1 := s[i]
		if v1 == '(' || v1 == '{' || v1 == '[' {
			stack = append(stack, v1)
		} else {
			n := len(stack)
			if n == 0 {
				return false
			}
			v2 := stack[n-1]
			stack = stack[:n-1]
			if v1 == ')' && v2 != '(' {
				return false
			}
			if v1 == '}' && v2 != '{' {
				return false
			}
			if v1 == ']' && v2 != '[' {
				return false
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

func isValid2(s string) bool {
	res := true
	s = matching(s)
	if len(s) > 0 {
		res = false
	}
	return res
}

func matching(s string) string {
	mapRightThese := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	mapLeftThese := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	if len(s)%2 > 0 {
		fmt.Println(s)
		return s
	}
	for i, v := range s {
		if i == 0 {
			continue
		}

		if _, ok := mapLeftThese[v]; ok {
			continue
		}

		if v2, ok := mapRightThese[v]; ok {
			if v2 == rune(s[i-1]) {
				return matching(s[:i-1] + s[i+1:])
			} else {
				return s
			}
		}
	}
	return s
}
