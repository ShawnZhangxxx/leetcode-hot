/*
__author__ = 'robin-luo'
__date__ = '2023/03/27 10:17'
*/

package main

import "fmt"

func main()  {
     res :=	climbStairs2(10)
     fmt.Println(res)
}

func climbStairs2(n int) int {
	dp := make([]int,n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}



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