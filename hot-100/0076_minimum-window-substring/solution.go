/*
__author__ = 'robin-luo'
__date__ = '2023/06/26 15:03'
*/

package main

func minWindow(s, t string) string {
	if len(s) < len(t) {
		return ""
	}

	m := make([]int, 256)
	for i := 0; i < len(t); i++ {
		m[t[i]]++
	}

	l, count, max, res := 0, len(t), len(s)+1, ""
	for i := 0; i < len(s); i++ {
		m[s[i]]--
		if m[s[i]] >= 0 {
			count--
		}

		for l < i && m[s[l]] < 0 {
			m[s[l]]++
			l++
		}

		if count == 0 && max > i-l+1 {
			max = i - l + 1
			res = s[l : i+1]
		}
	}

	return res
}
