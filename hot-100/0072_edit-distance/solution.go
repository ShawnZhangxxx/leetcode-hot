/*
__author__ = 'robin-luo'
__date__ = '2023/03/27 10:22'
*/

package main

import (
	"fmt"
)

func main()  {
	res:= minDistance2("intention","execution")
	fmt.Println(res)
}

func minDistance2(a string,b string) int {
	m := len(a)
	n := len(b)
	var dp = make([][]int,m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int,n+1)
	}
	if m*n == 0 { //有一个为0
		return m+n
	}

	for i := 0; i < m+1; i++ {
		dp[i][0] = i
	}
	for i := 0; i < n+1; i++ {
		dp[0][i] = i
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			dp[i][j] = min(dp[i-1][j-1],min(dp[i-1][j],dp[i][j-1])) + 1
			if  a[i-1] == b[j-1] { //表示前a的前i个字符 和 b的前j的  两个字符串的匹配状态.
				dp[i][j] = dp[i][j] - 1
			}
		}
	}

	return  dp[m][n]
}



func minDistance(word1, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 && n == 0 || word1 == word2 {
		return 0
	}

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for i := 0; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j])+1)
			}
		}
	}

	return dp[m][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
