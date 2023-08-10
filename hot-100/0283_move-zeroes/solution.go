/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-24 10:03:28
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-24 10:15:57
 * @Description:
 */

package main

func moveZeroes(nums []int) {
	i, j := 0, 0
	for j < len(nums) {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}

		j++
	}
}
