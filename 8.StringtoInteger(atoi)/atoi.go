package main

import (
	"fmt"
)

func main() {
	fmt.Println(myAtoi(" x"))
	fmt.Println(myAtoi("223"))
	fmt.Println(myAtoi("-223"))
	fmt.Println(myAtoi("+-223"))
	fmt.Println(myAtoi("2+-223"))
	fmt.Println(myAtoi("000011223"))
}

func myAtoi(str string) int {
	// + 43
	// - 45
	// 0 rune 48
	// 9 rune 57
	// " " 32
	//fmt.Println(" "[0])
	const INTMAX = int(^uint32(0) >> 1)
	const INTMIN = ^INTMAX

	isFirstSpace := true
	isFlag := true
	flag := 1
	res := 0
	for _, v := range str {
		if isFirstSpace == true && v == 32 {
			isFirstSpace = true
		} else {
			isFirstSpace = false
		}
		if isFirstSpace == true {
			continue
		}

		if isFlag == true && v == 45 {
			flag = -1
			isFlag = false
			continue
		} else if isFlag == true && v == 43 {
			flag = 1
			isFlag = false
			continue
		}

		if v > 47 && v < 58 {
			//fmt.Println(int(v))
			if flag == 1 && res*flag > INTMAX/10 || (res*flag == INTMAX/10 && v > 55) {
				return INTMAX
			}

			if flag == -1 && res*flag < INTMIN/10 || (res*flag == INTMIN/10 && v > 56) {
				return INTMIN
			}
			res = res*10 + int(v) - 48
		} else {
			break
		}
		isFlag = false
	}

	res = res * flag

	return res
}
