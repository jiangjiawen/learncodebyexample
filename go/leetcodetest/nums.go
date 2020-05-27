package main

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

// func minSubArray(s int, nums []int) int {
// 	midSave := []int{}
// 	var sum, min int

// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] >= s {
// 			return 1
// 		}

// 		sum += nums[i]
// 		midSave = append(midSave, nums[i])
// 		if sum >= s {
// 			for {
// 				if sum-B[0] < s {
// 					break
// 				}
// 				sum -= B[0]
// 				B = B[1:]
// 			}
// 			if min == 0 || min > len(B) {
// 				min = len(B)
// 			}
// 		}
// 	}
// 	return min
// }

func minStartValue(nums []int) int {
	var valueNeg = 0
	var compareFlage = 0
	for i := 0; i < len(nums); i++ {
		valueNeg += nums[i]
		if valueNeg < compareFlage {
			compareFlage = valueNeg
		}
	}
	if compareFlage >= 0 {
		return 1
	} else {
		return -compareFlage + 1
	}
}

func generate(numRows int) [][]int {
	yh := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		temp := make([]int, i+1)
		if i == 0 {
			temp[0] = 1
		}
		if i == 1 {
			temp[0] = 1
			temp[1] = 1
		} else {
			temp[0] = 1
			temp[i] = 1
			for j := 1; j < i; j++ {
				temp[j] = yh[i-1][j] + yh[i-1][j-1]
			}
		}
		yh[i] = temp
	}
	return yh
}

//面试51
func reversePairs(nums []int) int {
	var cnts = 0
	mergeTwo(nums, 0, len(nums)-1, []int{}, &cnts)
	return cnts
}

func merge(nums []int, start, mid, end int, temp []int, cnts *int) {
	i, j := start, mid+1
	for i <= mid && j <= end {
		if nums[i] <= nums[j] {
			temp = append(temp, nums[i])
			i++
		} else {
			*cnts = *cnts + mid - i + 1
			temp = append(temp, nums[j])
			j++
		}
	}
	for i <= mid {
		temp = append(temp, nums[i])
		i++
	}
	for j <= end {
		temp = append(temp, nums[j])
		j++
	}
	for i := 0; i < len(temp); i++ {
		nums[start+i] = temp[i]
	}
	temp = []int{}
}

func mergeTwo(nums []int, start, end int, temp []int, cnts *int) {
	if start >= end {
		return
	}
	mid := (start + end) >> 1
	mergeTwo(nums, start, mid, temp, cnts)
	mergeTwo(nums, mid+1, end, temp, cnts)
	merge(nums, start, mid, end, temp, cnts)
}

func main() {
	// fmt.Println(mySqrt(8))
	// fmt.Println(perfectnum(0))
	// fmt.Println(minStartValue([]int{-3, 2, -3, 4, 2}))
	// fmt.Println(generate(5))
	reversePairs([]int{1, 3, 2, 3, 1})
}
