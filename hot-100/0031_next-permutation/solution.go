/*
__author__ = 'robin-luo'
__date__ = '2023/03/10 15:18'
*/

package main

func nextPermutation(nums []int) {
	left := len(nums) - 2
	for 0 <= left && nums[left] >= nums[left+1] {
		left--
	}

	reverse(nums, left+1)
	if left == -1 {
		return
	}

	right := search(nums, left+1, nums[left])
	nums[left], nums[right] = nums[right], nums[left]
}

func reverse(nums []int, l int) {
	r := len(nums) - 1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func search(nums []int, l, target int) int {
	r := len(nums) - 1
	l--
	for l+1 < r {
		mid := l + (r-l)/2
		if target < nums[mid] {
			r = mid
		} else {
			l = mid
		}
	}

	return r
}
