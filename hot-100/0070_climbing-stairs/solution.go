/*
__author__ = 'robin-luo'
__date__ = '2023/03/27 10:17'
*/

package main

func climbStairs(n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = -1
	}

	dp[0], dp[1] = 1, 1
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}