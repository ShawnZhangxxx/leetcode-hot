/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-27 11:05:53
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-27 11:13:08
 * @Description:
 */

package main

func countSubstrings(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		count += expandAroundCenter(s, i, i)
		count += expandAroundCenter(s, i, i+1)
	}

	return count
}

func expandAroundCenter(s string, l, r int) int {
	count := 0
	for l >= 0 && r < len(s) && s[l] == s[r] {
		count++
		l--
		r++
	}

	return count
}
