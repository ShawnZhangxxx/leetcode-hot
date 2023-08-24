/*
__author__ = 'robin-luo'
__date__ = '2023/03/14 15:56'
*/

package main

import "fmt"

func permute(nums []int) [][]int {
	var (
		track []int
		res   [][]int
	)

	backTracking(nums, track, &res)
	return res
}
func permute2(nums []int) [][]int {
	var (
		track []int
		res   [][]int
	)

	backTracking2(nums, track, &res)
	return res
}
func main()  {
	res := permute2([]int{1,2,3})
	fmt.Println(res)
}

// track要放入结果的数组  nums要遍历的数组, res结果
func  backTracking2(nums ,track []int,res *[][]int)  {
	if  len(nums) == 0 {
		*res = append(*res,append([]int(nil),track...))//
	}
	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		track = append(track,cur)
		nums = append(nums[:i], nums[i+1:]...)
		backTracking2(nums,track,res)
		nums = append(nums[:i], append([]int{cur}, nums[i:]...)...)
		track = track[:len(track)-1]
	}
}


func backTracking(nums, track []int, res *[][]int) {
	if len(nums) == 0 {
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
	}
	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		track = append(track, cur)
		nums = append(nums[:i], nums[i+1:]...)
		backTracking(nums, track, res)
		nums = append(nums[:i], append([]int{cur}, nums[i:]...)...)
		track = track[:len(track)-1]
	}
}
