package main

import "fmt"

func main() {

	fmt.Println(commonChars([]string{"bella", "label", "roller"}))

	fmt.Println(commonChars([]string{"cool", "lock", "cook"}))

}

func commonChars(A []string) []string {
	res := []string{}
	minCountRune := map[rune]int{}

	for i := 0; i < len(A); i++ {
		countRune := map[rune]int{}
		for _, v := range A[i] {
			if vv, ok := countRune[v]; ok {
				countRune[v] = vv + 1
			} else {
				countRune[v] = 1
			}
		}

		if i == 0 {
			minCountRune = countRune
		} else {
			for rr := rune("a"[0]); rr <= rune("z"[0]); rr++ {
				if vv2, ok := countRune[rr]; ok {
					if vv3, ok2 := minCountRune[rr]; ok2 {
						if vv2 < vv3 {
							minCountRune[rr] = vv2
						}
					}
				} else {
					delete(minCountRune, rr)
				}
			}
		}
	}

	for k, v := range minCountRune {
		if v > 0 {
			for i := 0; i < v; i++ {
				res = append(res, string(k))
			}
		}
	}

	return res
}
