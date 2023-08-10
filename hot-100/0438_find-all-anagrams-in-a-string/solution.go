/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-26 10:45:28
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-26 10:59:18
 * @Description:
 */

package main

func findAnagrams(s, p string) []int {
	res := make([]int, 0)
	pCount := [26]int{}
	for _, c := range p {
		pCount[c-'a']++
	}

	l, r := 0, 0
	windowCharCount := [26]int{}
	for r < len(s) {
		windowCharCount[s[r]-'a']++
		if r-l+1 == len(p) {
			if isAnagram(windowCharCount, pCount) {
				res = append(res, l)
			}

			windowCharCount[s[l]-'a']--
			l++
		}

		r++
	}

	return res
}

func isAnagram(windowCharCount, pCount [26]int) bool {
	for i := 0; i < 26; i++ {
		if windowCharCount[i] != pCount[i] {
			return false
		}
	}

	return true
}
