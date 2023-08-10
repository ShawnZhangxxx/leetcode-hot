/*
__author__ = 'robin-luo'
__date__ = '2023/03/15 10:55'
*/

package main

import "math"

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	cover := 0
	for i := 0; i <= cover; i++ {
		cover = int(math.Max(float64(cover), float64(i+nums[i])))
		if cover >= len(nums)-1 {
			return true
		}
	}

	return false
}
