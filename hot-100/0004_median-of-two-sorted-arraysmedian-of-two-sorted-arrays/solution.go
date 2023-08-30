/*
__author__ = 'robin-luo'
__date__ = '2023/03/07 16:02'
*/

package main

import "math"

func main() {

}

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	left := (m + n + 1) / 2
	right := (m + n + 2) / 2
	return float64(findKth(nums1, nums2, 0, 0, left)+findKth(nums1, nums2, 0, 0, right)) / 2.0
}

func findKth(nums1, nums2 []int, i, j, k int) int {
	if i >= len(nums1) { //不可能> ,只有等于,被上一轮排空了
		return nums2[j+k-1]
	}

	if j >= len(nums2) { //
		return nums1[i+k-1]
	}

	if k == 1 {//如果只找一个
		return min(nums1[i], nums2[j])
	}

	var midVal1, midVal2 int
	if i+k/2-1 < len(nums1) {     //找+2/k后的中间值,如果超过最大长度直接砍 另一个数组
		midVal1 = nums1[i+k/2-1]
	} else {
		midVal1 = math.MaxInt64
	}

	if j+k/2-1 < len(nums2) {   //找+2/k后的中间值,如果超过最大长度直接砍 另一个数组
		midVal2 = nums2[j+k/2-1]
	} else {
		midVal2 = math.MaxInt64
	}

	if midVal1 < midVal2 {//2大,砍1 k/2
		return findKth(nums1, nums2, i+k/2, j, k-k/2)
	} else {
		return findKth(nums1, nums2, i, j+k/2, k-k/2)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
