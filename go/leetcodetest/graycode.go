package main

import "fmt"

//1238
func circularPermutation(n int, start int) []int {
	graycode := []int{}
	startloc := 0
	for i := 0; i < 1<<n; i++ {
		graycode = append(graycode, i^i>>1)
		if i^i>>1 == start {
			startloc = i
		}
	}

	return append(graycode[startloc:], graycode[:startloc]...)
}

func main() {
	fmt.Println(circularPermutation(2, 1))
}
