package main

import "fmt"

// x^5 = x^1*x^4
// x^13 = x^1*x^4*x^8
func Solution(x, n int) int {
	if n < 0 {
		return 0
	}

	var res = 1
	for n > 0 {
		if n%2 == 1 {
			res *= x
		}
		x *= x
		n /= 2
	}
	return res
}

func main() {
	fmt.Println(Solution(3, 3))
	fmt.Println(Solution(2, 10))
}
