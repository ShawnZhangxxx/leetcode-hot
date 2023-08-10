/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-26 15:51:33
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-27 09:37:27
 * @Description:
 */

package main

func findTargetSumWays(nums []int, target int) int {
	count := 0
	findTarget(nums, target, 0, 0, &count)
	return count
}

func findTarget(nums []int, target, index int, currentSum int, count *int) {
	if index == len(nums) {
		if currentSum == target {
			*count++
		}

		return
	}

	findTarget(nums, target, index+1, currentSum+nums[index], count)
	findTarget(nums, target, index+1, currentSum-nums[index], count)
}
