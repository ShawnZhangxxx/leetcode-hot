/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-03 14:30:24
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-04 09:54:04
 * @Description:
 */

package main

import "fmt"

//这道题
func main()  {
	res := wordBreak2("leetcode",[]string{"leet","code"})
	fmt.Println(res)
}

func wordBreak2(s string,wordDict []string)bool  {
	dp := make([]bool,len(s)+1)
	fmt.Println(dp)
	dp[0] = true //dp[i]表示前 0-i个字符是不是已经是全单词组成,j分割指针当 dp[j,i] && dp[j]为true 结束当前节点的循环,为true

	for i := 0; i < len(s); i++ { //从0开始
		for j := i; j >= 0 ; j-- {
			if dp[j] && indexOf(s[j:i+1],wordDict) > -1 {
				dp[i+1] = true
				break
			}
		}
	}
	fmt.Println(dp)
	return dp[len(s)]
}

func indexOf(s string,wordDict []string)int  {
	for i := 0; i < len(wordDict); i++ {
		if s == wordDict[i] {
			return i
		}
	}
	return -1
}















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
