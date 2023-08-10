/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-03 13:56:06
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-03 14:24:36
 * @Description:
 */

package main

import "sort"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)
	longestStreak, curStreak := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			if nums[i] == nums[i-1]+1 {
				curStreak++
			} else {
				longestStreak = max(longestStreak, curStreak)
				curStreak = 1
			}
		}
	}

	return max(longestStreak, curStreak)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
