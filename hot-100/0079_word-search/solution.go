/*
__author__ = 'robin-luo'
__date__ = '2023/06/26 16:58'
*/

package main

func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[0] {
			if board[i][j] == word[0] {
				if dfs(i, j, board, word[1:]) {
					return true
				}
			}
		}
	}

	return false
}

func dfs(x, y int, board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}

	// up
	if x > 0 && board[x-1][y] == word[0] {
		temp := board[x][y]
		board[x][y] = '#'
		if dfs(x-1, y, board, word[1:]) {
			return true
		}

		board[x][y] = temp
	}

	// down
	if x < len(board)-1 && board[x+1][y] == word[0] {
		temp := board[x][y]
		board[x][y] = '#'
		if dfs(x+1, y, board, word[1:]) {
			return true
		}

		board[x][y] = temp
	}

	// left
	if y > 0 && board[x][y-1] == word[0] {
		temp := board[x][y]
		board[x][y] = '#'
		if dfs(x, y-1, board, word[1:]) {
			return true
		}

		board[x][y] = temp
	}

	// right
	if y < len(board[0])-1 && board[x][y+1] == word[0] {
		temp := board[x][y]
		board[x][y] = '#'
		if dfs(x, y+1, board, word[1:]) {
			return true
		}

		board[x][y] = temp
	}

	return false
}
