package main

import "fmt"

func main() {
	target := 9
	nums := []int{2, 7, 11, 18}
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
	lmap := make(map[int]int)
	var i, j int
	for k, v := range nums {
		if target <= v || 2*v == target {
			continue
		}
		lmap[v] = k
		fmt.Println(k)
	}
	//fmt.Printf("%v\n", lmap)
	for k := range lmap {
		i = k
		j = target - i
		if _, ok := lmap[j]; ok {
			break
		}
	}
	return []int{lmap[i], lmap[j]}
}
