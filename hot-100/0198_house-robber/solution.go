/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-10 09:38:34
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-10 09:46:22
 * @Description:
 */

package main

func rob(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	if size == 1 {
		return nums[0]
	}

	dp := make([]int, size)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < size; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}

	return dp[size-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
