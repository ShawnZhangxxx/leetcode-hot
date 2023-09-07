///*
// * @Author: lawtech0902 584563542@qq.com
// * @Date: 2023-07-14 10:44:28
// * @LastEditors: lawtech0902 584563542@qq.com
// * @LastEditTime: 2023-07-18 10:37:00
// * @Description:
// */
//
package main
//
//import (
//	"math"
//)
//
//func maximalSquare(matrix [][]byte) int {
//	m := len(matrix)
//	n := len(matrix[0])
//	if m == 0 || n == 0 {
//		return 0
//	}
//
//	dp := make([][]int, m)
//	for i := range dp {
//		dp[i] = make([]int, n)
//	}
//
//	var maxSize int
//	for i := 0; i < m; i++ {
//		for j := 0; j < n; j++ {
//			if matrix[i][j] == '0' {
//				continue
//			}
//
//			if i == 0 || j == 0 {
//				dp[i][j] = 1
//			} else {
//				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
//			}
//
//			if maxSize < dp[i][j] {
//				maxSize = dp[i][j]
//			}
//		}
//	}
//
//	return maxSize * maxSize
//}
//
//func min(nums ...int) int {
//	min := math.MaxInt32
//	for _, num := range nums {
//		if num < min {
//			min = num
//		}
//	}
//
//	return min
//}
