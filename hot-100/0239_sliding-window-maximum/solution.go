/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-20 17:13:06
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-21 10:34:35
 * @Description:
 */

package main

import "fmt"

func main()  {
	res := maxSlidingWindow2([]int{1,3,-1,-3,5,3,6,7},3)
	fmt.Println(res)
}

func maxSlidingWindow2( nums []int,k int)[]int  {
	var q []int
	push := func(i int) { //从右向左扣,扣掉所有比当前元素小的数
		for len(q) > 0 && q[len(q)-1 ] < nums[i] {
			q = q[:len(q)-1]
		}
		q = append(q,i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	ans := make([]int,1,len(nums)-k+1)
	ans[0] = nums[q[0]]
	for i := k; i < len(nums); i++ {
		push(i)
		for j := 0; j < len(q); j++ {//从左往右扣,扣掉那些在窗口左边的越界的
			if q[j] < i-k+1 {
				q = q[1:]
			}
		}
		fmt.Println(q) //这个队列是单调递减的
		ans = append(ans,nums[q[0]])
	}
	return ans

}


























func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	res := make([]int, 0, n-k+1)
	// 双端队列维护最大值索引
	deque := make([]int, 0, k)

	for i := 0; i < n; i++ {
		// 如果队列中的第一个元素已经不在当前窗口内，将其从队列中删除
		for len(deque) > 0 && deque[0] < i-k+1 {
			deque = deque[1:]
		}

		// 如果队列中的元素比当前元素小，那么它们不可能是窗口中的最大值，因此将它们从队列中删除
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}

		// 将当前元素放入队列中
		deque = append(deque, i)

		// 当窗口的大小达到 k 时，队列的第一个元素就是当前窗口的最大值，将其添加到结果数组中
		if i >= k-1 {
			res = append(res, nums[deque[0]])
		}
	}

	return res
}
