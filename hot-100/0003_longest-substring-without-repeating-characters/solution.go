/*
__author__ = 'robin-luo'
__date__ = '2023/03/07 15:02'
*/

package main

import "fmt"

func main() {
	maxLen := lengthOfLongestSubstring2("aaaaa")
	fmt.Println(maxLen)
}

//滑动窗口  双指针 最长不重复串
func lengthOfLongestSubstring2(s string) int {

	m := map[string]int{}
	// for _, v := range s {  //不能便利字符串 ，变成rune型了
	// 	fmt.Println(v)
	// 	// if _, ok := m[v]; ok {

	// 	// }
	// }
	var fast int = 0
	var slow int = 0
	maxLen := 0
	for i, v := range s {

		if _, ok := m[string(v)]; ok { //匹配到
			if maxLen < (i - slow + 1) {
				maxLen = i - slow + 1
			}
			slow = m[string(s[i])] + 1
			for j := slow; j < m[string(s[i])]; j++ {
				delete(m, string(s[j]))
			}

		} else { //没匹配到

		}
		m[string(v)] = i
		fast = i
	}
	if maxLen < (fast - slow + 1) {
		maxLen = fast - slow + 1
	}
	return maxLen
}

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
