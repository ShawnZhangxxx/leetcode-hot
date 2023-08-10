/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-21 11:09:30
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-24 09:59:54
 * @Description:
 */

package main

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
