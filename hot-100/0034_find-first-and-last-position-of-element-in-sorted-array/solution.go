/*
__author__ = 'robin-luo'
__date__ = '2023/03/13 09:42'
*/

package main

import (
	"fmt"
	"sort"
)

func main()  {
	res := searchRange3([]int{1,2,3,3,4,5,5,10},0)
	fmt.Println(res)
}
//
func searchRange3(nums []int, target int) []int {
	leftmost := sort.SearchInts(nums, target)//查找左边界的二分,
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := sort.SearchInts(nums, target + 1) - 1
	return []int{leftmost, rightmost}
}

//func searchRange2(nums []int,target int) []int  {
//	l,r := 0,len(nums) -1
//	mid := 0
//	for l<=r {
//		mid = (l+r)/2
//		if nums[mid] < target { // 删 左边
//
//		}
//	}
//}

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			res := []int{0, 0}
			if nums[l] == target {
				res[0] = l
			}

			if nums[r] == target {
				res[1] = r
			}

			for i := mid; i < r+1; i++ { //这第儿代码可能导致不是log(n)的复杂度,如果重复数过多的话
				if nums[i] != target {
					res[1] = i - 1
					break
				}
			}

			for i := mid; i > l-1; i-- {
				if nums[i] != target {
					res[0] = i + 1
					break
				}
			}

			return res
		}
	}

	return []int{-1, -1}
}
