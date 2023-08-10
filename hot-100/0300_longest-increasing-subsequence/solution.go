/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-24 10:44:42
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-24 11:06:17
 * @Description:
 */

package main

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	maxLen := 0

	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		maxLen = max(maxLen, dp[i])
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
