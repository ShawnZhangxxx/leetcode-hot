/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-24 16:07:46
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-24 16:41:54
 * @Description:
 */

package main

import (
	"fmt"
)
/*

*/

func main()  {
	res := maxProfit([]int{1,2,3,0,2})
	fmt.Println(res)
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	//关键是0,1,2状态 0:当前持有股票可能是今天买的或者之前就有
	//1:不持有 ,明天冷冻期就是 下一天不能买
	f0, f1, f2 := -prices[0], 0, 0
	for i := 1; i < n; i++ {
		newf0 := max(f0, f2 - prices[i])  //当天持有股票 1,昨天也持有 2.今天买的,昨天的2状态不处在冷冻期状态
		newf1 := f0 + prices[i]  //不持有 处在冷冻期  今天卖股票,
		newf2 := max(f1, f2) //不持有,不处在冷冻期
		f0, f1, f2 = newf0, newf1, newf2
	}
	return max(f1, f2)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

