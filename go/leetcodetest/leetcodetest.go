package main

import "fmt"

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	r := x
	for r > x/r {
		r = (r + x/r) / 2
	}
	return int(r)
}

func perfectnum(n int) bool {
	if n == 0 {
		return true
	} else {
		var sum int = 0
		for i := 1; i*i <= n; i++ {

			if n%i == 0 {
				sum += i
				if i*i != n {
					sum += (n / i)
				}
			}

		}
		return sum-n-n == 0
	}

}

func main() {
	fmt.Println(mySqrt(8))
	fmt.Println(perfectnum(0))
}
