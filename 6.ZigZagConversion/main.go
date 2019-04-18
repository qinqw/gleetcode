package main

import (
	"fmt"
)

func main() {
	s := "skdiejjidmwlqwii23jdmw"
	row := 3
	fmt.Println(convert(s, row))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var ret string
	n := len(s)
	cycleLen := 2*numRows - 2

	for i := 0; i < numRows; i++ {
		for j := 0; j+i < n; j += cycleLen {
			ret += string(s[j+i])
			if i != 0 && i != numRows-1 && j+cycleLen-i < n {
				ret += string(s[j+cycleLen-i])
			}

		}
	}
	return ret
}
