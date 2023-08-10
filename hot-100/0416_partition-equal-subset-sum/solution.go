/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-26 10:13:34
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-26 10:25:26
 * @Description:
 */

package main

func canPartition(nums []int) bool {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	if totalSum%2 != 0 {
		return false
	}

	target := totalSum / 2
	dp := make([]bool, target+1)
	dp[0] = true
	for _, num := range nums {
		for j := target; j >= num; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}

	return dp[target]
}
