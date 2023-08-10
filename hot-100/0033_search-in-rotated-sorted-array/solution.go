/*
__author__ = 'robin-luo'
__date__ = '2023/03/10 15:50'
*/

package main

func search(nums []int, target int) int {
	size := len(nums)
	l, r := 0, size-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[l] {
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target < nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}
