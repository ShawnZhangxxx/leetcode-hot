/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-07 15:10:35
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-07 16:16:57
 * @Description:
 */

package main

func majorityElement(nums []int) int {
	major, count := 0, 0
	for _, num := range nums {
		if count == 0 {
			major = num
		}

		if major == num {
			count++
		} else {
			count--
		}
	}

	return major
}
