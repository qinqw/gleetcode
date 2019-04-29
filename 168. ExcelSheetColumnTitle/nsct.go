package main

import "fmt"

func main() {
	fmt.Println(convertToTitle(260))
	fmt.Println(convertToTitle(2600))
}

func convertToTitle(n int) string {
	mapI := map[int]string{
		1: "A", 2: "B", 3: "C", 4: "D", 5: "E",
		6: "F", 7: "G", 8: "H", 9: "I", 10: "J",
		11: "K", 12: "L", 13: "M", 14: "N", 15: "O",
		16: "P", 17: "Q", 18: "R", 19: "S", 20: "T",
		21: "U", 22: "V", 23: "W", 24: "X", 25: "Y",
		26: "Z",
	}
	res := ""
	//res := make([]rune, 0)
	for n > 0 {
		i := n % 26
		if i == 0 {
			i = 26
			n = (n - 26) / 26
		} else {
			n = n / 26
		}
		res = mapI[i] + res
		//fmt.Printf("i,n,s:%d,%d,%v\n", i, n, res)
	}
	return res
}
