/*
__author__ = 'robin-luo'
__date__ = '2023/06/26 14:47'
*/

package main

func sortColors(nums []int) {
	// 一次扫描，3指针法
	p0, p1 := 0, 0
	p2 := len(nums) - 1
	for p1 <= p2 {
		if nums[p1] == 0 {
			nums[p0], nums[p1] = nums[p1], nums[p0]
			p0++
			p1++
		} else if nums[p1] == 2 {
			nums[p1], nums[p2] = nums[p2], nums[p1]
			p2--
		} else {
			p1++
		}
	}
}
