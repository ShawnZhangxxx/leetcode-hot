/*
__author__ = 'robin-luo'
__date__ = '2023/03/17 10:49'
*/

package main

import "fmt"

func main()  {
	//[[1,3],[2,6],[8,10],[15,18]]
	res := uniquePaths2(2,3)
	fmt.Println(res)
}

func uniquePaths2(m, n int) int {
	// dp: dp[i][j] = dp[i-1][j] + dp[i][j-1]
	//dp := [m][n]int{}
	dp := make([][]int,m)//二维数组初始化 只能这样
	for i, _ := range dp {
		dp[i] = make([]int,n)
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	fmt.Println(dp)
	return dp[m-1][n-1]
}

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
