/*
__author__ = 'robin-luo'
__date__ = '2023/06/26 16:58'
*/

package main

import "fmt"

func main()  {
		res := exist2([][]byte{
			{'a','b','c','e'},
			{'s','f','c','s'},
			{'d','a','c','b'},
		},"abcccad")
		fmt.Println(res)
}

func exist2(board [][]byte,word string)bool  {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
		  if  dfs2(board,word,0,0,0) {
		  	return true
		  }
		}
	}
	return false

}

func dfs2(board [][]byte , word string,i int,j int,k int) bool  {
	if i >= len(board) || i<0 || j >= len(board[0]) || j<0 || board[i][j] != word[k] {
		return false
	}
	//能执行到这儿说明找到对的相应的字符
	if k == len(word) - 1 {
		return true
	}

	board[i][j] = '#' //防止回到起点
	res := dfs2(board,word,i+1,j,k+1)  ||
		dfs2(board,word,i-1,j,k+1) ||
		dfs2(board,word,i,j+1,k+1) ||
		dfs2(board,word,i,j-1,k+1)

	return res

}





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
