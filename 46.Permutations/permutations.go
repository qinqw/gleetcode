package main

import "fmt"

func main() {
	//fmt.Println(permute([]int{1}))
	fmt.Println(permute([]int{1, 2, 3, 4}))

}

func permute(nums []int) [][]int {
	var t1, t2 [][]int
	//res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		t2 = make([][]int, 0)
		//t1 = res
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
	//res = t1
	return t1
}
