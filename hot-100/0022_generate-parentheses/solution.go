/*
__author__ = 'robin-luo'
__date__ = '2023/03/09 16:45'
*/

package main

import "fmt"

func main() {
	// ans := []string{}
	// backtrack(&ans, "", 0, 0, 5)
	// fmt.Println(ans)

	ans2 := dfs2("", 2, 2)

	fmt.Println(ans2)
}

//方法2 深度优先算法
func dfs2(prefix string, nLeft int, nRight int) []string {
	if nLeft < 0 || nRight < 0 {
		return nil
	}

	if nLeft == 0 && nRight == 0 {
		return []string{prefix}
	}

	if nLeft > nRight {
		return nil
	}

	if nLeft == nRight {
		return dfs2(prefix+"(", nLeft-1, nRight)
	}

	var res []string
	r := dfs2(prefix+"(", nLeft-1, nRight)
	res = append(res, r...)
	r = dfs2(prefix+")", nLeft, nRight-1)
	res = append(res, r...)
	return res

}

//1.回溯法 最后一个cur = 可以省略，要传递数组指针，或者搞个全局变量 ，回溯法关键
func backtrack(ans *[]string, cur string, open int, close int, max int) {
	if len(cur) == max*2 {
		*ans = append(*ans, cur)
		return
	}
	if open < max {
		cur = cur + string('(')
		backtrack(ans, cur, open+1, close, max)
		cur = cur[:len(cur)-1] //回溯法关键代码
	}
	if open > close {
		cur = cur + string(')')
		backtrack(ans, cur, open, close+1, max)
		// cur = cur[:len(cur)-1]
	}
}

func generateParenthesis(n int) []string {
	return dfs("", n, n)
}

func dfs(prefix string, nLeft, nRight int) []string {
	if nLeft < 0 || nRight < 0 {
		return nil
	}

	if nLeft == 0 && nRight == 0 {
		return []string{prefix}
	}

	if nLeft > nRight {
		return nil
	}

	if nLeft == nRight {
		return dfs(prefix+"(", nLeft-1, nRight)
	}

	var res []string
	r := dfs(prefix+"(", nLeft-1, nRight)
	res = append(res, r...)
	r = dfs(prefix+")", nLeft, nRight-1)
	res = append(res, r...)
	return res
}
