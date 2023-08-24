/*
__author__ = 'robin-luo'
__date__ = '2023/06/26 15:18'
*/

package main

import "fmt"

func main()  {
	res:= subsets2([]int{1,2,3})
	fmt.Println(res)
}

func subsets2(nums []int)(ans [][]int) {
	n := len(nums)
	for mask:=0;mask < 1<<n;mask++	{
		set := []int{}
		for i,v := range nums {
			if mask >> i & 1>0 {
				set = append(set,v)
			}
		}
		ans = append(ans,append([]int(nil),set...))
	}
	return
}


func subsets(nums []int) [][]int {
	var (
		track []int
		res   [][]int
	)
	backTracking(nums, track, 0, &res)
	return res
}

func backTracking(nums, track []int, startIndex int, res *[][]int) {
	temp := make([]int, len(track))
	copy(temp, track)
	*res = append(*res, temp)
	for i := startIndex; i < len(nums); i++ {
		track = append(track, nums[i])
		backTracking(nums, track, i+1, res)
		track = track[:len(track)-1]
	}
}
