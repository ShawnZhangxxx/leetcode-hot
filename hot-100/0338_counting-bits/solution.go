/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-25 11:11:56
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-25 13:47:09
 * @Description:
 */

package main

func countBits(n int) []int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i/2] + i%2
	}

	return dp
}
