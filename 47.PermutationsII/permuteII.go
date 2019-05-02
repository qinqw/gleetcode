package main

import "fmt"

func main() {
	//fmt.Println(permute([]int{1, 2, 3, 4}))
	fmt.Println(permute([]int{1, 1, 3}))
}

func permuteUnique(nums []int) [][]int {
	//tmp := permute()
	return permute(nums)
}

func permute(nums []int) [][]int {
	var t1, t2, res [][]int
	for i := 0; i < len(nums); i++ {
		t2 = make([][]int, 0)
		if i == 0 {
			tt := append(make([]int, 0), nums[i])
			t1 = append(t1, tt)
			continue
		}
		for j := 0; j < len(t1); j++ {
			cur := t1[j]
			for k := 0; k <= len(cur); k++ {
				tt := make([]int, 0)
				if k == 0 {
					tt = append(append(make([]int, 0), nums[i]), cur...)
				} else if k == len(cur) {
					tt = append(cur, nums[i])
				} else {
					tt2 := []int{}
					tt2 = append(tt2, cur...)
					tt2 = append(tt2[:k], nums[i])
					tt = append(tt2, cur[k:]...)
				}
				t2 = append(t2, tt)
			}
		}
		t1 = t2
	}
	tmp := make(map[string]byte)
	for i := 0; i < len(t1); i++ {
		s := fmt.Sprintf("%v", t1[i])
		if _, ok := tmp[s]; !ok {
			res = append(res, t1[i])
			tmp[s] = 1
		}
	}
	return res
}
