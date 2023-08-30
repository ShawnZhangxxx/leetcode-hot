/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-06-28 17:13:16
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-06-28 17:22:27
 * @Description:
 */

package main

import "fmt"

func main()  {
	res := maxProfit2([]int{7,1,5,3,6,4})
	fmt.Println(res)
}

//besttime 其实只有两个时机,不是多个买卖时机
func maxProfit2(prices []int) int {
	low := prices[0]
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < low {
			low = prices[i]
		}
		maxProfit = max(maxProfit,prices[i]-low)
	}
	return maxProfit
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	low := prices[0]
	maxProfit := 0
	for _, price := range prices {
		if price < low {
			low = price
		}

		maxProfit = max(maxProfit, price-low)
	}

	return maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
