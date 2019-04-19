package main

import (
	"fmt"
)

func main() {
	//fmt.Println(romanToInt(""))
	fmt.Println(romanToInt("X"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("XIX"))
	fmt.Println(romanToInt("MCMXCIX"))
}

func romanToInt(s string) int {
	res := 0
	rImap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	l := len(s)
	for i := 0; i < l-1; i++ {
		if rImap[s[i]] < rImap[s[i+1]] {
			res -= rImap[s[i]]
			res += rImap[s[i+1]]
			i++
		} else {
			res += rImap[s[i+1]]
		}
	}
	if l > 0 {
		res += rImap[s[l-1]]
	}
	return res
}

func romanToInt2(s string) int {
	var sum int
	mp := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	for i := 0; i < len(s)-1; i++ {
		if mp[string(s[i])] < mp[string(s[i+1])] {
			sum -= mp[string(s[i])]
		} else {
			sum += mp[string(s[i])]
		}
	}
	return sum + mp[string(s[len(s)-1])]
}
