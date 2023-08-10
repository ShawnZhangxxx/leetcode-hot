/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-06-28 11:00:12
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-06-28 13:59:14
 * @Description:
 */

package main

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	return dp[n]
}
