/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-26 11:03:39
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-26 13:44:27
 * @Description:
 */

package main

func findDisapearedNumbers(nums []int) []int {
	res := make([]int, 0)
	for _, num := range nums {
		index := abs(num) - 1
		if nums[index] > 0 {
			nums[index] = -nums[index]
		}
	}

	for i, num := range nums {
		if num > 0 {
			res = append(res, i+1)
		}
	}

	return res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
