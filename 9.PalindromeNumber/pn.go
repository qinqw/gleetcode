package main

func main() {

}

func isPalindrome(x int) bool {
	res := true
	if x < 0 {
		return false
	}
	if x >= 0 && x < 10 {
		return true
	}
	tmp := []int{}
	for x > 0 {
		m := x % 10
		x = x / 10
		tmp = append(tmp, m)
		if x < 10 {
			tmp = append(tmp, x)
			break
		}
	}
	l := len(tmp)
	for i := 0; i < l; i++ {
		if 2*i <= l+1 {
			if tmp[i] != tmp[l-1-i] {
				res = false
				break
			}
		} else {
			break
		}

	}
	return res
}
