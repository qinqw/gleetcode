package main

func main() {

}

func reverse(x int) int {
	const INTMAX = int(^uint32(0) >> 1)
	const INTMIN = ^INTMAX
	rev := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		if rev > INTMAX/10 || (rev == INTMAX/10 && pop > 7) {
			return 0
		}
		if rev < INTMIN/10 || (rev == INTMIN/10 && pop < -8) {
			return 0
		}
		rev = rev*10 + pop
	}
	return rev
}
