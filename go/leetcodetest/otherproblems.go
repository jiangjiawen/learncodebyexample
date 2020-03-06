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

//849
// func maxDistToClosest(seats []int) int {
//     maxDis :=0
//     numofzero :=0
//     for i:=0;i<len(seats);i++{
//         if seats[i]==0{
//             numofzero += 1
//             if i==(numofzero-1){
//                 maxDis=numofzero
//             }
//         }
//         if seats[i]==1{
//             dish := (numofzero+numofzero%2)/2
//             numofzero = 0
//             if dish > maxDis {
//                 maxDis=dish
//             }
//         }
//     }
//     if numofzero > maxDis{
//         maxDis = numofzero
//     }
//     return maxDis
// }
func maxDistToClosest(seats []int) int {
	size := len(seats)
	maxDis := 0
	// e 代表了连续空位的个数
	// 当连续空位两边都有人的时候，maxDis = (e+e%2)/2
	// 如果有一边没人的话，      maxDis = e
	e := 0
	for i := 0; i < size; i++ {
		if e == i {
			// 说明 seats[0:i] 全是 0
			maxDis = e
		} else {
			maxDis = max(maxDis, (e+e%2)/2)
		}
		if seats[i] == 1 {
			e = 0
		} else {
			e++
		}
	}

	// 当 seats[size-1]==0 的时候
	// e 最后的值，有可能 > maxDis
	return max(maxDis, e)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
	seats := []int{1,0,0,0,1}
	fmt.Println(maxDistToClosest(seats))
}
