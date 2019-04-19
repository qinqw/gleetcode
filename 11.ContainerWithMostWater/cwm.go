package main

import "fmt"

func main() {

	//fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 8, 8, 3, 7}))

	// fmt.Println(maxArea([]int{4, 4, 5, 8, 5, 8, 3, 6, 5}))
	// fmt.Println(maxArea([]int{5, 4, 5, 2, 5, 8, 3, 3, 4}))
	// fmt.Println(maxArea([]int{5, 4, 4, 2, 5, 4, 5, 4, 5}))
	//fmt.Println(maxArea([]int{5, 5, 5, 5, 5, 5, 5, 5, 5}))

	//fmt.Println(maxArea([]int{9, 6, 14, 11, 2, 2, 4, 9, 3, 8}))
	fmt.Println(maxArea([]int{4, 4, 2, 11, 0, 11, 5, 11, 13, 8}))
}

func maxArea(height []int) int {
	var h int
	res := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			if height[i] < height[j] {
				h = height[i]
			} else {
				h = height[j]
			}
			if res < h*(j-i) {
				res = h * (j - i)
			}
		}
	}
	return res
}

func maxArea2(height []int) int {
	res := 0
	top := [2]int{}
	top2 := [2]int{}
	for i := 0; i < len(height); i++ {
		if height[i] > height[top[0]] && height[i] >= height[top[1]] {
			top[0], top[1] = top[1], top[0]
			top[1] = i
		} else if height[i] > height[top[0]] && height[i] < height[top[1]] {
			top[0] = i
		} else if height[i] == height[top[0]] && height[i] == height[top[1]] {
			top[1] = i
		}
		//fmt.Printf("i, top : %d, %v\n", i, top)
	}
	if top[0] > top[1] {
		top[0], top[1] = top[1], top[0]
	}
	//fmt.Println(top)
	top2 = top
	temp := 0
	for j := top[1]; j < len(height); j++ {
		var h int
		if height[top[0]] > height[j] {
			h = height[j]
		} else {
			h = height[top[0]]
		}

		t := h * (j - top[0])
		if temp < t {
			temp = t
			top2[1] = j
		}
	}

	//fmt.Println(top2)
	for k := top[0]; k >= 0; k-- {
		var h int
		if height[top2[1]] > height[k] {
			h = height[k]
		} else {
			h = height[top2[1]]
		}
		t := h * (top2[1] - k)
		if temp < t {
			temp = t
			top2[0] = k
		}
		// fmt.Printf("temp, t : %d, %d\n", t, h)
		// fmt.Printf("i, top : %d, %v\n", k, top2)
	}

	fmt.Println(top2)
	var h int
	if height[top2[1]] > height[top2[0]] {
		h = height[top2[0]]
	} else {
		h = height[top2[1]]
	}

	res = h * (top2[1] - top2[0])

	return res
}
