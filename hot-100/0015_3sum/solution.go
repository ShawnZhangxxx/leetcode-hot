/*
__author__ = 'robin-luo'
__date__ = '2023/03/08 10:46'
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	res := threeSum2([]int{1, 0, -1, 3, 4, -3})
	fmt.Println(res)
}
//1.排序 2.遍历一遍,从改数两个边界往里面挤
func threeSum2(nums []int) [][]int {

	var res [][]int
	sort.Ints(nums)
	var l, r int

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return res
		}
		l = i + 1
		r = len(nums) - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if l > 0 && nums[l] == nums[l-1] {
				l++
				continue
			}
			if r < len(nums)-1 && nums[r] == nums[r+1] {
				r--
				continue
			}
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
			} else {
				if sum < 0 {
					l++
				} else {
					r--
				}
			}
		}

	}

	return res
}

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	size := len(nums)
	if size <= 2 {
		return res
	}

	for i := range nums[:size-2] {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := nums[i] * -1
		l, r := i+1, size-1
		for l < r {
			if nums[l]+nums[r] == target {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}

				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if nums[l]+nums[r] > target {
				r--
			} else {
				l++
			}
		}
	}

	return res
}
