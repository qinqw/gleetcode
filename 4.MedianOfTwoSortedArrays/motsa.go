package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{3, 5}
	nums2 := []int{4}
	f := findMedianSortedArrays(nums1, nums2)
	fmt.Println(f)

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	f = findMedianSortedArrays(nums1, nums2)
	fmt.Println(f)

}

//MedianOfTwoSortedArrays is find the median from two sorted arrays
func MedianOfTwoSortedArrays(nums1, nums2 []int) float64 {
	res := 0.0
	//lnNums := append(nums1, nums2)

	return res
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	res := 0.0
	lnNums := append(nums1, nums2...)
	sort.Ints(lnNums)
	lenN := len(lnNums)
	fmt.Println(lenN)

	if lenN%2 == 1 {
		index := int(lenN / 2)
		res = float64(lnNums[index])
	} else {
		index := int(lenN / 2)
		res = (float64(lnNums[index-1]) + float64(lnNums[index])) / 2
	}

	return res
}
