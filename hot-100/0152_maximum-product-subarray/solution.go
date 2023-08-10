/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-05 10:59:11
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-05 16:29:48
 * @Description:
 */

package main

import "math"

func maxProduct(nums []int) int {
	res := math.MinInt64
	iMax, iMin := 1, 1
	for _, num := range nums {
		if num < 0 {
			iMax, iMin = iMin, iMax
		}

		iMax = max(iMax*num, num)
		iMin = min(iMin*num, num)
		res = max(res, iMax)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
