/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-21 11:09:30
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-24 09:59:54
 * @Description:
 */

package main

import (
	"fmt"
	"math"
)

func main()  {
	res:= perfectSquare2(198)
	fmt.Println(res)
}

func perfectSquare2(num int)int  {
	dp := make([]int,num+1)
	//dp[0] = 0
	for i := 1; i <= num; i++ {
		minNum := math.MaxInt64
		for j := 1; j * j <= i; j++ {
			minNum = min(minNum,dp[i-j*j])
		}
		dp[i] = minNum + 1

	}
	return dp[num]
}
















func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
