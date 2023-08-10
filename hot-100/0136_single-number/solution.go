/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-03 14:25:13
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-03 14:29:48
 * @Description:
 */

package main

func singleNumber(nums []int) int {
	a := 0
	for _, num := range nums {
		a ^= num
	}

	return a
}
