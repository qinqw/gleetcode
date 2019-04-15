package main

import (
	"fmt"
	"testing"
)

func TestTwoSum(test *testing.T) {
	target := 9
	nums := []int{2, 7, 11, 18}
	fmt.Printf("%v", TwoSum(nums, target))
}
