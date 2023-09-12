/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-24 16:07:46
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-24 16:41:54
 * @Description:
 */

package main

import "math"
/*
sold：
前一天hold，当日卖出股票

hold： 可由两个情况转换来
前一天hold，当日rest
前一天rest，当日买入股票变为hold

rest：
前一天sold，当日必须rest
前一天rest，当日继续rest
*/
func maxProfit(prices []int) int {
	sold, rest, hold := 0, 0, math.MinInt32
	for _, p := range prices {
		preSold := sold
		sold = hold + p
		hold = max(hold, rest-p)
		rest = max(rest, preSold)
	}

	return max(sold, rest)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
