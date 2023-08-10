/*
__author__ = 'robin-luo'
__date__ = '2023/03/17 10:49'
*/

package main

func uniquePaths(m, n int) int {
	// dp: dp[i][j] = dp[i-1][j] + dp[i][j-1]
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = 0
		}
	}

	dp[0][0] = 1
	for i := range dp {
		for j := range dp[i] {
			if i+1 < m {
				dp[i+1][j] += dp[i][j]
			}

			if j+1 < n {
				dp[i][j+1] += dp[i][j]
			}
		}
	}

	return dp[m-1][n-1]
}
