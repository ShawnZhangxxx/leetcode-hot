/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-21 10:43:18
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-21 10:58:33
 * @Description:
 */

package main

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}

	n := len(matrix[0])

	// 从矩阵右上角开始查找
	row := 0
	col := n - 1
	for row < m && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			row++
		} else {
			col--
		}
	}

	return false
}
