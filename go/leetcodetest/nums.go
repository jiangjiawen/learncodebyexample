package main

import "fmt"

//x的平方根
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

//完美数
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

func minSubArray(s int, nums []int) int {
	midSave := []int{}
	var sum, min int

	for i := 0; i < len(nums); i++ {
		if nums[i] >= s {
			return 1
		}

		sum += nums[i]
		midSave = append(midSave, nums[i])
		if sum >= s {
			for {
				if sum-B[0] < s {
					break
				}
				sum -= B[0]
				B = B[1:]
			}
			if min == 0 || min > len(B) {
				min = len(B)
			}
		}
	}
	return min
}

func main() {
	fmt.Println(mySqrt(8))
	fmt.Println(perfectnum(0))
}
