package main

import (
"fmt"
)

func main() {
	n := 4 // 设置N皇后问题的N值
	board := make([][]bool, n)
	for i := range board {
		board[i] = make([]bool, n)
	}

	solveNQueens(board, 0)
}

func solveNQueens(board [][]bool, col int) {
	if col >= len(board) {
		printBoard(board)
		return
	}

	for row := 0; row < len(board); row++ {
		if isSafe(board, row, col) {
			board[row][col] = true
			solveNQueens(board, col+1)
			board[row][col] = false
		}
	}
}

func isSafe(board [][]bool, row, col int) bool {
	// 检查列上是否有冲突
	for i := 0; i < len(board); i++ {
		if board[i][col] {
			return false
		}
	}

	// 检查左上方是否有冲突
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] {
			return false
		}
	}

	// 检查右上方是否有冲突
	for i, j := row-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] {
			return false
		}
	}

	return true
}

func printBoard(board [][]bool) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] {
				fmt.Print("Q ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}