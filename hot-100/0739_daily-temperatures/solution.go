/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-27 11:13:47
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-27 13:38:30
 * @Description:
 */

package main

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[idx] = i - idx
		}

		stack = append(stack, i)
	}

	return res
}
