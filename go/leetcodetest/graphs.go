package main

import (
	"fmt"
	"sort"
)

//332
//https://leetcode-cn.com/problems/reconstruct-itinerary/solution/custerxue-xi-bi-ji-dfs-tu-de-ou-la-lu-jing-by-cust/
func findItinerary(tickets [][]string) []string {
	m := make(map[string][]string, len(tickets)+1)
	var ans []string

	for _, t := range tickets {
		m[t[0]] = append(m[t[0]], t[1])
	}

	fmt.Println(m)

	for k := range m {
		sort.Strings(m[k])
	}

	fmt.Println(m)

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

	row1 := []string{"JFK", "SFO"}
	row2 := []string{"JFK", "ATL"}
	row3 := []string{"SFO", "ATL"}
	row4 := []string{"ATL", "JFK"}
	row5 := []string{"ATL", "SFO"}
	tickets = append(tickets, row1)
	tickets = append(tickets, row2)
	tickets = append(tickets, row3)
	tickets = append(tickets, row4)
	tickets = append(tickets, row5)

	// row1 := []string{"JFK", "KUL"}
	// row2 := []string{"JFK", "NRT"}
	// row3 := []string{"NRT", "JFK"}
	// tickets = append(tickets, row1)
	// tickets = append(tickets, row2)
	// tickets = append(tickets, row3)
	fmt.Println(findItinerary(tickets))
}
