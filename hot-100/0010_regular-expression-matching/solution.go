/*
__author__ = 'robin-luo'
__date__ = '2023/03/07 17:14'
*/

package main

import "fmt"

func main()  {
	flg := isMatch2("abcdefaaa","abc.*")
	fmt.Println(flg)
}
//matches 函数用于判断字符串 s 的第 i 个字符是否能够与正则表达式 p 的第 j 个字符匹配。
//
//使用二维数组 f，其中 f[i][j] 表示字符串 s 的前 i 个字符是否能够与正则表达式 p 的前 j 个字符匹配。
//
//初始化 f[0][0] = true，表示空字符串能够与空正则表达式匹配。
//
//然后使用双重循环来填充数组 f。对于每个 f[i][j]，有两种情况：
//
//a. 若 p[j-1] 为 *，则可以考虑使用或不使用 *，即 f[i][j] 可以由 f[i][j-2] 或者如果当前字符匹配，由 f[i-1][j] 得到。
//
//b. 若 p[j-1] 不为 *，并且当前字符匹配，则 f[i][j] 可以由 f[i-1][j-1] 得到。
//
//最后，返回 f[m][n]，其中 m 为字符串 s 的长度，n 为正则表达式 p 的长度，即判断整个字符串 s 是否与正则表达式 p 匹配。
func isMatch2(s, p string) bool {
	m,n := len(s),len(p)
	matches := func (i,j int)bool {
		if i== 0 {
			return false
		}
		if p[j-1] == '.' {
			return  true
		}
		return s[i-1] == p[j-1]
	}
	f := make([][]bool,m+1) //为了0,0只有一个元素的时候,所以才+1
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool,n+1)
	}
	f[0][0] = true //表示空字符串能与空正则表达式匹配
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i][j-2]
				if matches(i, j-1){ //前一个字符与 i 匹配上了
					f[i][j] = f[i][j] || f[i-1][j]
				}
			}else if matches(i,j) {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}

	return f[m][n]

}



func isMatch(s, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true
	for j := 1; j < n && dp[0][j-1]; j += 2 {
		if p[j] == '*' {
			dp[0][j+1] = true
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if p[j] == '.' || p[j] == s[i] {
				dp[i+1][j+1] = dp[i][j]
			} else if p[j] == '*' {
				if p[j-1] != s[i] && p[j-1] != '.' {
					dp[i+1][j+1] = dp[i+1][j-1]
				} else {
					dp[i+1][j+1] = dp[i+1][j-1] || dp[i+1][j] || dp[i][j+1]
				}
			}
		}
	}

	return dp[m][n]
}
