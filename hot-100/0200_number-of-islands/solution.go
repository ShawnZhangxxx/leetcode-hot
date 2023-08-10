/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-10 10:02:46
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-10 10:11:36
 * @Description:
 */

package main

func numIslands(grid [][]byte) int {
	var (
		res int
		dfs func(i, j int)
	)

	dfs = func(i, j int) {
		if 0 <= i && i < len(grid) && 0 <= j && j < len(grid[i]) && grid[i][j] == '1' {
			grid[i][j] = '0'
			dfs(i+1, j) // right
			dfs(i-1, j) // left
			dfs(i, j+1) // down
			dfs(i, j-1) // up
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				res++
				dfs(i, j)
			}
		}
	}

	return res
}
