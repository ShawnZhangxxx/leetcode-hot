/*
__author__ = 'robin-luo'
__date__ = '2023/03/18 19:45'
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	res := minPathSum2([][]int{{1,2,3},{4,5,6}})
	fmt.Println(res)
}




func minPathSum2(grid [][]int) int {
	m,n := len(grid),len(grid[0])
	dp := make([][]int,m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int,n)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	for j := 1; j < m; j++ {
		dp[j][0] = dp[j-1][0] + grid[j][0]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = int(math.Min(float64(dp[i-1][j]),float64(dp[i][j-1]))) + grid[i][j]
		}
	}
	fmt.Println(dp)
	return dp[m-1][n-1]
}


func minPathSum(grid [][]int) int {
	if grid == nil {
		return 0
	}

	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = 0
		}
	}

	dp[0][0] = grid[0][0]
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = int(math.Min(float64(dp[i][j-1]), float64(dp[i-1][j]))) + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}
