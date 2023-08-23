/*
__author__ = 'robin-luo'
__date__ = '2023/06/26 14:47'
*/

package main

import "fmt"

//0 1 2
func main()  {
	nums := []int{0,1,1,1,1,2,2,0,1,2}
	 sortColors2(nums)
	fmt.Println(nums)
}

//其实就是单指针 两个边界
func sortColors2(nums []int)  {
	l,r := 0,len(nums) -1
	p := 0
	for p<=r {
		if nums[p] == 0 {
			nums[l],nums[p] = nums[p],nums[l]
			p++
			l++
		}else if nums[p] == 2 {
			nums[p],nums[r] = nums[r], nums[p]
			r--
		} else {
			p++
		}
	}
}



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
