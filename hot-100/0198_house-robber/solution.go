/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-10 09:38:34
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-10 09:46:22
 * @Description:
 */

package main

import "fmt"

func main()  {
	//res := rob2([]int{2,7,9,3,1})
	res := rob2([]int{1,2,3,1})
	fmt.Println(res)

}

func rob2(nums []int)  int{
	dp := make([]int,len(nums) + 1)
	dp[0] = 0
	dp[1] = nums[0]
	dp[2] = nums[0]
	for i := 3; i <= len(nums); i++ {
		dp[i] = max(dp[i-2],dp[i-3]) + nums[i-1]
	}
	fmt.Println(dp)
	return max(dp[len(nums)],dp[len(nums) -1])

}


func rob(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	if size == 1 {
		return nums[0]
	}

	dp := make([]int, size)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < size; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}

	return dp[size-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
