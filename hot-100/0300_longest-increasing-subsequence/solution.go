/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-24 10:44:42
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-24 11:06:17
 * @Description:
 */

package main

import "fmt"

func main()  {
	res := lengthOfLIS2([]int{10,9,2,5,3,7,101,18})
	fmt.Println(res)
}

//这是连续递增子序列
func lengthOfLIS2(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	maxLen := 0
	dp[0] =1
	for i := 1; i < n; i++ {
		dp[i] = 1
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1]+1
		}
		maxLen = max(maxLen, dp[i])
	}
	return maxLen
}

//这才是
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
