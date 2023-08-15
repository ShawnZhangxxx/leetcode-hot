/*
__author__ = 'robin-luo'
__date__ = '2023/03/10 15:18'
*/

package main

import "fmt"

func main()  {
	nums := nextPermutation2([]int{5,6,7,3,2,1})
	fmt.Println(nums)
}

//找下一个增长字符序列的切片
func  nextPermutation2(nums []int)[]int  {
	if len(nums) <= 1 {
		return nil
	}
	i,j,k := len(nums)-2 ,len(nums)-1,len(nums)-1

	//1.第一步 找到1个降序的
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}
	if i>= 0 { //找到比i位置刚好大的,用来交换,正好所谓的自然序列增长一点
		for nums[i] >= nums[k] { //可以二分法,找到k
			k--
		}
		nums[i],nums[k] = nums[k],nums[i]
	}
	//尾部剩余颠倒顺序  s++,e--不行只能这么写
	for s,e := j,len(nums) -1; s < e; s,e = s+1,e-1  {
		nums[s],nums[e] = nums[e],nums[s]
	}

	return nums

}

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
