package main

import "fmt"

//888 索引处的字符串
func decodeAtIndex(S string, K int) string {
	var count int = 0
	for i := 0; i < len(S); i++ {
		a := []rune(S)[i]
		if a <= '9' && a >= '0' {
			if ((int)(a-'0') * count) >= K {
				if K%count == 0 {
					return decodeAtIndex(S, K-count)
				} else {
					return decodeAtIndex(S, K%count)
				}
			} else {
				count = (int)(a-'0') * count
			}
		} else {
			count++
		}
		if count == K {
			return string(a)
		}
	}
	return ""
}

// 44
//https://leetcode-cn.com/problems/wildcard-matching/solution/dong-tai-gui-hua-dai-zhu-shi-by-tangweiqun/
func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(p)+1)
	}

	dp[len(s)][len(p)] = true
	for i := len(s); i >= 0; i-- {
		for j := len(p); j >= 0; j-- {
			if i == len(s) && j == len(p) {
				continue
			}
			if j != len(p) && ((i != len(s)) && (s[i] == p[j] || p[j] == '?')) {
				dp[i][j] = dp[i+1][j+1]
			} else if (j != len(p)) && p[j] == '*' {
				dp[i][j] = ((i != len(s)) && (dp[i+1][j]) || dp[i][j+1])
			} else {
				dp[i][j] = false
			}
		}
	}

	return dp[0][0]
}

func main() {
	S := "leet2code3"
	K := 10
	fmt.Println(decodeAtIndex(S, K))
	fmt.Println(isMatch("abcde", "*a*e"))
}
