package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(multiply("0", "12"))
	fmt.Println(multiply("", "12"))
	fmt.Println(multiply("1", "12"))
	fmt.Println(multiply("1888381881212", "1212988182128182818218281828812"))
}

func multiply(num1 string, num2 string) string {
	l1 := len(num1)
	l2 := len(num2)
	if l1 < 1 || l2 < 1 || num1 == "0" || num2 == "0" {
		return "0"
	}
	l := l1 + l2
	tmp := make([]int, l)
	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			t1, _ := strconv.Atoi(string(num1[l1-1-i]))
			t2, _ := strconv.Atoi(string(num2[l2-1-j]))
			t := t1 * t2
			tmp[l-1-i-j] += t
		}
	}

	res := ""
	tmpI := 0
	for k := 0; k < l; k++ {
		if tmp[l-1-k] > 0 {
			tmpI = l - 1 - k
		}
		t1 := tmp[l-1-k] % 10
		t2 := tmp[l-1-k] / 10
		res = strconv.Itoa(t1) + res
		if t2 > 0 {
			tmp[l-1-k-1] += t2
		}
	}
	//fmt.Printf("tmpI,tmp:%d,%s\n", tmpI, res)
	res = res[tmpI:]
	return res
}
