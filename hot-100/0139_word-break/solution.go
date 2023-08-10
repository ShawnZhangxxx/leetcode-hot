/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-03 14:30:24
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-04 09:54:04
 * @Description:
 */

package main

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 0; i < n; i++ {
		for j := i; j >= 0; j-- {
			if dp[j] == true && inWordDict(s[j:i+1], wordDict) {
				dp[i+1] = true
			}
		}
	}

	return dp[n]
}

func inWordDict(s string, wordDict []string) bool {
	for _, word := range wordDict {
		if word == s {
			return true
		}
	}

	return false
}
