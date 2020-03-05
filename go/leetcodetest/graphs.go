package main

import (
	"fmt"
	"sort"
)

//332
// func RemoveIndex(s []int, index int) []int {
// 	return append(s[:index], s[index+1:]...)
// }

// func contains(arr []int, str int) bool {
// 	for _, a := range arr {
// 		if a == str {
// 			return true
// 		}
// 	}
// 	return false
// }

// func findItinerary(tickets [][]string) []string {
// 	mem := make(map[int][]int)
// 	start := []int{}
// 	savelist := []int{}
// 	for i, ticket := range tickets {
// 		for j := 0; j < len(tickets); j++ {
// 			if tickets[j][0] == ticket[1] && j != i {
// 				mem[i] = append(mem[i], j)
// 			}
// 		}
// 		if ticket[0] == "JFK" && len(tickets) == 1 {
// 			start = append(start, i)
// 		} else if ticket[0] == "JFK" && len(mem[i]) != 0 {
// 			start = append(start, i)
// 		}
// 		savelist = append(savelist, i)
// 	}

// 	startnode := start[0]
// 	for i := 1; i < len(start); i++ {
// 		if strings.Compare(tickets[startnode][1], tickets[start[i]][1]) > 0 {
// 			startnode = start[i]
// 		} else {
// 			continue
// 		}
// 	}
// 	// fmt.Println(mem)
// 	lastlist := RemoveIndex(savelist, startnode)
// 	res := []int{}
// 	res = append(res, startnode)
// 	for len(lastlist) != 0 {
// 		mostloc := -1
// 		if len(mem[res[len(res)-1]]) == 0 {
// 			mostloc = -1
// 			break
// 		} else if len(mem[res[len(res)-1]]) == 1 {
// 			if contains(lastlist, mem[res[len(res)-1]][0]) {
// 				mostloc = mem[res[len(res)-1]][0]
// 			} else {
// 				break
// 			}

// 		} else {
// 			si := 1
// 			if contains(lastlist, mem[res[len(res)-1]][0]) {
// 				mostloc = mem[res[len(res)-1]][0]
// 			} else {
// 				mostloc = mem[res[len(res)-1]][1]
// 				si = 2
// 			}
// 			for j := si; j < len(mem[res[len(res)-1]]); j++ {
// 				if contains(lastlist, mem[res[len(res)-1]][j]) && strings.Compare(tickets[mostloc][0], tickets[mem[res[len(res)-1]][j]][0]) > 0 {
// 					mostloc = mem[res[len(res)-1]][j]
// 				}
// 			}
// 		}
// 		if mostloc != -1 {
// 			res = append(res, mostloc)
// 			tmplist := []int{}
// 			for k := range lastlist {
// 				if lastlist[k] != mostloc {
// 					tmplist = append(tmplist, lastlist[k])
// 				}
// 			}
// 			lastlist = tmplist
// 		}
// 	}

// 	resstring := []string{}

// 	for i := range res {
// 		resstring = append(resstring, tickets[res[i]][0])
// 	}
// 	resstring = append(resstring, tickets[res[len(res)-1]][1])
// 	return resstring
// }

func findItinerary(tickets [][]string) []string {
	m := make(map[string][]string, len(tickets)+1)
	var ans []string

	for _, t := range tickets {
		m[t[0]] = append(m[t[0]], t[1])
	}

	for k := range m {
		sort.Strings(m[k])
	}

	DFS("JFK", m, &ans)

	// revers ans array
	i, j := 0, len(ans)-1
	for i < j {
		ans[i], ans[j] = ans[j], ans[i]
		i++
		j--
	}
	return ans
}

func DFS(start string, m map[string][]string, ans *[]string) {
	for len(m[start]) > 0 {
		cur := m[start][0]
		m[start] = m[start][1:]
		DFS(cur, m, ans)
	}

	*ans = append(*ans, start)
}

func main() {
	var tickets [][]string

	// row1 := []string{"MUC", "LHR"}
	// row2 := []string{"JFK", "MUC"}
	// row3 := []string{"SFO", "SJC"}
	// row4 := []string{"LHR", "SFO"}
	// tickets = append(tickets, row1)
	// tickets = append(tickets, row2)
	// tickets = append(tickets, row3)
	// tickets = append(tickets, row4)

	// row1 := []string{"JFK", "SFO"}
	// row2 := []string{"JFK", "ATL"}
	// row3 := []string{"SFO", "ATL"}
	// row4 := []string{"ATL", "JFK"}
	// row5 := []string{"ATL", "SFO"}
	// tickets = append(tickets, row1)
	// tickets = append(tickets, row2)
	// tickets = append(tickets, row3)
	// tickets = append(tickets, row4)
	// tickets = append(tickets, row5)

	row1 := []string{"JFK", "KUL"}
	row2 := []string{"JFK", "NRT"}
	row3 := []string{"NRT", "JFK"}
	tickets = append(tickets, row1)
	tickets = append(tickets, row2)
	tickets = append(tickets, row3)
	fmt.Println(findItinerary(tickets))
}
