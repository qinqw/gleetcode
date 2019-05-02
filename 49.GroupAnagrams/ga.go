package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}

// type sByte []byte

// func (l sByte) Len() int           { return len(l) }
// func (l sByte) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
// func (l sByte) Less(i, j int) bool { return l[i] < l[j] }

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	tmap := make(map[string]int)
	for i := 0; i < len(strs); i++ {
		str := []byte(strs[i])
		sort.Slice(str, func(i, j int) bool {
			return str[i] < str[j]
		})
		index := string(str)
		if v, ok := tmap[index]; ok {
			res[v] = append(res[v], strs[i])
		} else {
			tmap[index] = len(tmap)
			res = append(res, []string{strs[i]})
		}
	}
	return res
}
