package main

import "fmt"

func main() {
	target := 9
	nums := []int{2, 7, 11, 18}
	fmt.Printf("%v", TwoSum(nums, target))

	target = 4
	nums = []int{2, 2, 11, 18}
	fmt.Printf("%v", TwoSum(nums, target))

	target = 6
	nums = []int{2, 3, 11, 3}
	fmt.Printf("%v", TwoSum(nums, target))

	target = 6
	nums = []int{3, 3, 11}
	fmt.Printf("%v", TwoSum(nums, target))

	target = 6
	nums = []int{3, 2, 4}
	fmt.Printf("%v", TwoSum(nums, target))

	target = 0
	nums = []int{-3, 2, 3, 5}
	fmt.Printf("%v", TwoSum(nums, target))
}

// TwoSum is
// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// Example:
// 	Given nums = [2, 7, 11, 15], target = 9,
// 	Because nums[0] + nums[1] = 2 + 7 = 9,
//  return [0, 1].
func TwoSum(nums []int, target int) []int {
	//var lmap map[int]int
	var res []int
	lmap := make(map[int]int)
	tmpA := make([]int, 0)
	var i, j int
	for k, v := range nums {
		if target == 2*v {
			tmpA = append(tmpA, k)
		} else {
			lmap[v] = k
		}

	}
	//fmt.Printf("%v\n", lmap)
	if len(tmpA) > 1 {
		res = []int{tmpA[0], tmpA[1]}
	} else {
		for k := range lmap {
			i = k
			j = target - i
			if _, ok := lmap[j]; ok {
				break
			}
		}
		res = []int{lmap[i], lmap[j]}
	}

	return res
}
