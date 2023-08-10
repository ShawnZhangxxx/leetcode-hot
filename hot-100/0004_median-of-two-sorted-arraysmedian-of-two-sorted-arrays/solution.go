/*
__author__ = 'robin-luo'
__date__ = '2023/03/07 16:02'
*/

package main

import (
	"fmt"
	"math"
)

func  main()  {
	k :=	findMedianSortedArrays2([]int{1,2,3,3,3},[]int{1,4,6,7})
	fmt.Println(k)
}
func findMedianSortedArrays2(nums1, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	var midValue float64
	sumLength := m+n
	if sumLength % 2 == 0 {
		midValue = 	float64(findKth2(nums1,nums2,0,0, sumLength / 2) + findKth2(nums1,nums2,0,0, sumLength / 2 + 1) )/ 2.0
	}else {
		midValue = 	float64(findKth2(nums1,nums2,0,0, sumLength / 2 + 1))
	}
	return  midValue
}

func findKth2(nums1, nums2 []int, i, j, k int) int {
	//-1目的是 起始数组下标是0, 目标k位置的个元素,在k-1下标的位置, k/2 在 k/2  -1 的位置 ,
	// k/2 -1 + k/2 -1 <= k - 2  (k-1个数字), 如果是 k/2 + k/2 是 可能是 k+1个数
	if i >= len(nums1) { // 当左边界溢出了, i=len(nums1)时就已经溢出了
		return nums2[j  + k - 1] //
	}

	if j >= len(nums1) {
		return nums1[j  + k - 1]
	}

	var midVal1,midVal2 int  //计算中位值 数组k/2位置值,如果剩余长度不够k/2 ,直接删另一个数组,将中位值提到最大
	if i+k/2 -1 < len(nums1) { //没溢出,len 与下标位比较
		midVal1 = nums1[i+k/2 -1]
	} else { //nums1 大小不够半个k了,一定删num2
		midVal1 = math.MaxUint32
	}

	if j+k/2-1 < len(nums2) { //不到num1 的右边界
		midVal2 = nums2[j+k/2 -1]
	} else { //nums1 大小不够半个k了,一定删num2
		midVal2 = math.MaxUint32
	}

	if midVal1 < midVal2 {
		return findKth(nums1,nums2,i+k/2,j,k-k/2)
	}else {
		return findKth(nums1,nums2,i,j+k/2,k-k/2)

	}
}




func findMedianSortedArrays(nums1, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	left := (m + n + 1) / 2
	right := (m + n + 2) / 2
	return float64(findKth(nums1, nums2, 0, 0, left)+findKth(nums1, nums2, 0, 0, right)) / 2.0 //这种计算量是别人的两倍
}

func findKth(nums1, nums2 []int, i, j, k int) int {
	if i >= len(nums1) {
		return nums2[j+k-1]
	}

	if j >= len(nums2) {
		return nums1[i+k-1]
	}

	if k == 1 {
		return min(nums1[i], nums2[j])
	}

	var midVal1, midVal2 int
	if i+k/2-1 < len(nums1) {
		midVal1 = nums1[i+k/2-1]
	} else {
		midVal1 = math.MaxInt64
	}

	if j+k/2-1 < len(nums2) {
		midVal2 = nums2[j+k/2-1]
	} else {
		midVal2 = math.MaxInt64
	}

	if midVal1 < midVal2 {
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
