/*
__author__ = 'robin-luo'
__date__ = '2023/03/08 10:46'
*/

package main

import "sort"

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
