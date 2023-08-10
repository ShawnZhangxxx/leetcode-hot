/*
__author__ = 'robin-luo'
__date__ = '2023/03/09 16:45'
*/

package main

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
