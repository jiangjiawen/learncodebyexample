package main

import (
	"fmt"
	"sort"
)

type boxType [][]int

func (b boxType) Len() int           { return len(b) }
func (b boxType) Less(i, j int) bool { return b[i][0] < b[j][0] }
func (b boxType) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

//0813
func pileBox(box [][]int) int {
	sort.Sort(boxType(box))
	dp := make([]int, len(box))
	for i := range box {
		dp[i] = 0
	}
	maxdepth := 0
	for i := range box {
		dp[i] = box[i][2]
		for j := 0; j < i; j++ {
			if box[j][0] < box[i][0] && box[j][1] < box[i][1] && box[j][2] < box[i][2] {
				if dp[i] > dp[j]+box[i][2] {
					dp[i] = dp[i]
				} else {
					dp[i] = dp[j] + box[i][2]
				}
			}
		}
		if maxdepth > dp[i] {
			continue
		} else {
			maxdepth = dp[i]
		}
	}
	return maxdepth
}

func main() {
	box := [][]int{}
	row1 := []int{1, 1, 1}
	row2 := []int{2, 2, 2}
	row3 := []int{3, 3, 3}
	box = append(box, row1)
	box = append(box, row2)
	box = append(box, row3)
	fmt.Println(pileBox(box))
}
