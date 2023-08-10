/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-24 16:53:04
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-24 17:18:13
 * @Description:
 */

package main

/*
dp + 分治 top->down
nums左右两侧+1，进行分治
*/
func maxCoins(nums []int) int {
	n := len(nums)
	nums = append([]int{1}, append(nums, 1)...)
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}

	for l := 1; l <= n; l++ {
		for i := 1; i <= n-l+1; i++ {
			j := i + l - 1
			for k := i; k <= j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k-1]+nums[i-1]*nums[k]*nums[j+1]+dp[k+1][j])
			}
		}
	}

	return dp[1][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
