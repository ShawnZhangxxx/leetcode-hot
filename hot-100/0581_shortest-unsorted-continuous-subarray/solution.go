/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-27 10:25:24
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-27 10:29:18
 * @Description:
 */

package main

import "math"

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	l, r := -1, -1
	maxNum, minNum := math.MinInt32, math.MaxInt32

	for i := 0; i < n; i++ {
		if nums[i] < maxNum {
			r = i
		} else {
			maxNum = nums[i]
		}
	}

	for i := n - 1; i >= 0; i-- {
		if nums[i] > minNum {
			l = i
		} else {
			minNum = nums[i]
		}
	}

	if r == -1 {
		return 0
	}

	return r - l + 1
}
