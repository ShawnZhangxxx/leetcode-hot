/*
__author__ = 'robin-luo'
__date__ = '2023/03/07 15:02'
*/

package main

import "fmt"

// 滑动窗口
func lengthOfLongestSubstring(s string) int {
	location := [128]int{}
	for i := range location {
		location[i] = -1
	}

	maxLen, left := 0, 0
	for i := 0; i < len(s); i++ {
		if location[s[i]] >= left {
			left = location[s[i]] + 1
		} else if i+1-left > maxLen {
			maxLen = i + 1 - left
		}

		location[s[i]] = i
	}

	return maxLen
}
func main() {
	maxLen := lengthOfLongestSubstring("12345654321")
	fmt.Println(maxLen)
}