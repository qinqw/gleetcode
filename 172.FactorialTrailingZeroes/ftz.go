package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(trailingZeroes(0))
	fmt.Println(trailingZeroes(0))

	fmt.Println(trailingZeroes(3))
	fmt.Println(trailingZeroes(-3))

	fmt.Println(trailingZeroes(-8))
	fmt.Println(trailingZeroes(8))

	fmt.Println(trailingZeroes(-11))
	fmt.Println(trailingZeroes(11))

	fmt.Println(trailingZeroes(-19))
	fmt.Println(trailingZeroes(19))

	fmt.Println(trailingZeroes(-100))
	fmt.Println(trailingZeroes(99))
}

func trailingZeroes(n int) int {
	n = int(math.Abs(float64(n)))
	f1 := 0
	for n > 5 {
		n = n / 5
		f1 += n
	}
	return f1
}
