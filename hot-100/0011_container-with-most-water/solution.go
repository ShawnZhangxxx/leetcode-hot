/*
__author__ = 'robin-luo'
__date__ = '2023/03/08 10:32'
*/

package main

import "math"

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	res, area := 0, 0
	for left < right {
		if height[left] < height[right] {
			area = height[left] * (right - left)
			left++
		} else {
			area = height[right] * (right - left)
			right--
		}

		res = int(math.Max(float64(res), float64(area)))
	}

	return res
}
