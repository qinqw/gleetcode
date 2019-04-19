package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(intToRoman(100))
	fmt.Println(intToRoman(5))
	fmt.Println(intToRoman(90))
	fmt.Println(intToRoman(1999))
	fmt.Println(intToRoman(3999))
}

func intToRoman(num int) string {
	iRMap := map[int]string{
		3000: "MMM",
		2000: "MM",
		1000: "M",
		900:  "CM",
		800:  "DCCC",
		700:  "DCC",
		600:  "DC",
		500:  "D",
		400:  "CD",
		300:  "CCC",
		200:  "CC",
		100:  "C",
		90:   "XC",
		80:   "LXXX",
		70:   "LXX",
		60:   "LX",
		50:   "L",
		40:   "XL",
		30:   "XXX",
		20:   "XX",
		10:   "X",
		9:    "IX",
		8:    "VIII",
		7:    "VII",
		6:    "VI",
		5:    "V",
		4:    "IV",
		3:    "III",
		2:    "II",
		1:    "I",
	}
	res := ""
	i := 0
	for num > 0 {
		t := num % 10
		if t > 0 {
			res = iRMap[t*int(math.Pow10(i))] + res
		}
		num = num / 10
		i++
	}

	return res

}
