/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-06-27 10:13:43
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-06-28 10:37:25
 * @Description:
 */
/*
__author__ = 'robin-luo'
__date__ = '2023/06/27 10:13'
*/

package main

func largestRectangleArea(heights []int) int {
	var (
		stack  []int
		i, res = 0, 0
	)

	for i < len(heights) {
		if len(stack) == 0 || heights[i] > heights[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			var width int
			if len(stack) == 0 {
				width = i
			} else {
				width = i - stack[len(stack)-1] - 1
			}

			res = max(res, width*heights[top])
			i--
		}

		i++
	}

	for len(stack) != 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		var width int
		if len(stack) == 0 {
			width = i
		} else {
			width = len(heights) - stack[len(stack)-1] - 1
		}

		res = max(res, width*heights[top])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
