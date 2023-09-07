/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-14 10:44:28
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-18 10:37:00
 * @Description:
 */

package main

import (
	"fmt"
	"math"
)

func main()  {
	res := maximalSquare([][]byte{
		{1,0,1,0,0},
		{1,0,1,1,1},
		{1,0,1,1,1},
		{1,0,1,1,1},
	})
	fmt.Println(res)
}
//如果这个点时正方形,左上方也是.
func maximalSquare(matrix [][]byte) int {
	dp := make([][]int,len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int,len(matrix[0]))
		dp[i][0] = int(matrix[i][0])
	}
	//初始化边界 也可以放到下面初始化,当i / j= 0时边界顺便初始化掉,因为用不上
	for i := 0; i < len(matrix[0]); i++ {
		dp[0][i] = int(matrix[0][i])
	}
	max := math.MinInt64
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 1 && matrix[i-1][j] == 1 && matrix[1][j-1] == 1  {
				dp[i][j] = dp[i-1][j-1]+1
				if max < dp[i][j] {
					max = dp[i][j]
				}
			}else if matrix[i][j] == 1 {
				dp[i][j] = 1
			}
		}
	}
	fmt.Println(dp)
	return  max * max
}

func min(nums ...int) int {
	min := math.MaxInt32
	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	return min
}
