/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-24 17:20:00
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-25 10:16:44
 * @Description:
 */

package main

func coinChange(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}

	// dp[i] 表示凑成金额 i 所需要的最少硬币数量
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		// 将 dp[0] 设置为 0，表示凑成金额为 0 不需要硬币
		// 将所有元素设置为一个较大的值，表示无穷大
		dp[i] = amount + 1
		for _, coin := range coins {
			if i >= coin {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	// 表示无法凑成金额amount
	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
