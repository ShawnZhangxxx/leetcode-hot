/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-20 15:51:09
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-20 17:07:11
 * @Description:
 */

package main

func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	res[0] = 1
	//去掉当前位置的前缀积
	for i := 1; i < n; i++ {
		res[i] = res[i-1] * nums[i-1]
	}
	//从右向左推,上面的*右侧的前缀积也是抛去当前元素
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return res
}
