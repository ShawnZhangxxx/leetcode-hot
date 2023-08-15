/*
__author__ = 'robin-luo'
__date__ = '2023/03/10 15:50'
*/

package main

import "fmt"

func main()  {
	lo := search2([]int{4,5,6,1,2,3},3)
	fmt.Println(lo)
}
//2023-8-15
func search2(num []int,target int)int  {

	l,r := 0,len(num) -1
	var mid int

	for l<=r  {  //这地儿是 <= ,
		mid = (l+r)/2
		if num[mid] == target {
			return mid
		}
		//总有一边是有序的
		if num[0] < num[mid] { //左边有序
			if target >= num[0] && target<num[mid] {
				r = mid -1
			} else  {
				l = mid + 1
			}
		}else { //右边有序
			if target > num[mid] && target<=num[r] {
				l = mid + 1
			} else  {
				r = mid - 1
			}
		}
	}
	return  -1
}















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
