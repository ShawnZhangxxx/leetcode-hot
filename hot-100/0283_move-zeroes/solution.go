/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-24 10:03:28
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-24 10:15:57
 * @Description:
 */

package main

import "fmt"

func main()  {
	nums := []int{0,1,2,3,0,0,1}
	moveZero2(nums)
	fmt.Println(nums)
}

func moveZero2( nums []int)  {
	z := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i],nums[z] = nums[z],nums[i]
			z++
		}
	}

}













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
