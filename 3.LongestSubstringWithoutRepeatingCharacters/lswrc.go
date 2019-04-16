package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	s := "abdiwhusssaa"
	fmt.Println(lengthOfLongestSubstring(s))

	s = "aaaaaa"
	fmt.Println(lengthOfLongestSubstring(s))

	s = "abcabcaa"
	fmt.Println(lengthOfLongestSubstring(s))

	s = "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))

	s = " "
	fmt.Println(lengthOfLongestSubstring(s))

	s = ""
	fmt.Println(lengthOfLongestSubstring(s))

	s = "abc"
	fmt.Println(lengthOfLongestSubstring(s))
	s = "ab"
	fmt.Println(lengthOfLongestSubstring(s))

	s = "dvdf"
	fmt.Println(lengthOfLongestSubstring(s))

	s = "asjrgapa"
	fmt.Println(lengthOfLongestSubstring(s))

	s = "ckilbkd"
	fmt.Println(lengthOfLongestSubstring(s))
}

// lengthOfLongestSubstring is find the substring
func lengthOfLongestSubstring(s string) int {
	res := 0
	tmpMaxLen := 0
	startIndex := 0
	runeOccrMap := [128]int{}
	startIndexMap := [128]int{}
	for i, c := range s {
		runeOccrMap[c] = runeOccrMap[c] + 1
		if runeOccrMap[c] == 1 {
			tmpMaxLen++
		} else if runeOccrMap[c] >= 2 {
			if startIndexMap[c] < startIndex {
				tmpMaxLen++
			} else {
				startIndex = startIndexMap[c]
				tmpMaxLen = i - startIndex
			}
		}
		if res < tmpMaxLen {
			res = tmpMaxLen
		}
		startIndexMap[c] = i
	}
	if res < tmpMaxLen {
		res = tmpMaxLen
	}
	return res
}

// lengthOfLongestSubstring is find the substring
func lengthOfLongestSubstring6(s string) int {
	res := 0
	tmpMaxLen := 0
	startIndex := 0
	runeOccrMap := [128]int{}
	startIndexMap := make(map[rune]int)
	for i, c := range s {
		runeOccrMap[c] = runeOccrMap[c] + 1
		if runeOccrMap[c] == 1 {
			tmpMaxLen++
		} else if runeOccrMap[c] >= 2 {
			if startIndexMap[c] < startIndex {
				tmpMaxLen++
			} else {
				startIndex = startIndexMap[c]
				tmpMaxLen = i - startIndex
			}
		}

		if res < tmpMaxLen {
			res = tmpMaxLen
		}

		startIndexMap[c] = i
		//fmt.Printf("%c", c)
	}
	//fmt.Printf("\n")
	if res < tmpMaxLen {
		res = tmpMaxLen
	}
	return res
}

// lengthOfLongestSubstring is find the substring
func lengthOfLongestSubstring5(s string) int {
	res := 0
	tmpMaxLen := 0
	startIndex := 0
	runeOccrMap := make(map[rune]int)
	startIndexMap := make(map[rune]int)
	for i, c := range s {
		//fmt.Printf("[%d] %c\n", i, c)
		if v2, ok2 := runeOccrMap[c]; ok2 {
			runeOccrMap[c] = v2 + 1
		} else {
			runeOccrMap[c] = 1
		}

		if runeOccrMap[c] == 1 {
			tmpMaxLen++
		} else if runeOccrMap[c] >= 2 {
			if startIndexMap[c] < startIndex {
				tmpMaxLen++
			} else {
				startIndex = startIndexMap[c]
				tmpMaxLen = i - startIndex
			}
		}

		if res < tmpMaxLen {
			res = tmpMaxLen
		}

		startIndexMap[c] = i
		fmt.Printf("%c", c)
	}
	fmt.Printf("\n")
	if res < tmpMaxLen {
		res = tmpMaxLen
	}
	return res
}

func lengthOfLongestSubstring3(s string) int {
	var wg sync.WaitGroup
	res := 0
	lenS := len(s)

	ch := make(chan (int), 5)
	for j := 0; j < lenS; j++ {
		wg.Add(1)
		go func(jj int) {
			//fmt.Printf("loop:%d\n", j+1)
			tlen := 0
			maxLen := 0
			tmpMap := make(map[byte]byte)
			for i := jj; i < lenS; i++ {
				//fmt.Println("i", i)
				if _, ok := tmpMap[s[i]]; ok {
					tmpMap = make(map[byte]byte)
					if tlen < maxLen {
						tlen = maxLen
					}
					maxLen = 1
					tmpMap[s[i]] = 0
				} else {
					tmpMap[s[i]] = 0
					maxLen++
				}
				if tlen < maxLen {
					tlen = maxLen
				}
			}
			ch <- tlen
			defer wg.Done()
		}(j)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for d := range ch {
		if res < d {
			//res = maxLen
			res = d
		}
		//fmt.Println(d)
		time.Sleep(100)
	}
	return res
}

// lengthOfLongestSubstring2 is find the substring
func lengthOfLongestSubstring2(s string) int {
	res := 0
	lenS := len(s)
	maxLen := 0
	tmpMap := make(map[byte]byte)
	for j := 0; j < lenS; j++ {
		maxLen = 0
		tmpMap = make(map[byte]byte)
		for i := j; i < lenS; i++ {
			//fmt.Println("i", i)
			if _, ok := tmpMap[s[i]]; ok {
				tmpMap = make(map[byte]byte)
				if res < maxLen {
					res = maxLen
				}
				maxLen = 1
				tmpMap[s[i]] = 0
			} else {
				tmpMap[s[i]] = 0
				maxLen++
			}
		}
		if res < maxLen {
			res = maxLen
		}
		// if res > lenS-j {
		// 	break
		// }
	}
	return res
}

// LongestSubstringWithoutRepeatingCharacters is find the substring
func LongestSubstringWithoutRepeatingCharacters(str string) int {
	res := 0
	maxLen := 0
	tmpMap := make(map[rune]byte)
	for _, c := range str {
		//fmt.Printf("[%d] %c\n", i, c)
		if _, ok := tmpMap[c]; ok {
			tmpMap := make(map[rune]byte)
			if res < maxLen {
				res = maxLen
			}
			maxLen = 1
			tmpMap[c] = 0
		} else {
			tmpMap[c] = 0
			maxLen++
		}
	}
	if res < maxLen {
		res = maxLen
	}
	return res
}
