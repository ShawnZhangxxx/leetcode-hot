/*
__author__ = 'robin-luo'
__date__ = '2023/03/07 17:02'
*/

package main

func main()  {
	p := longestPalindrome2("zbcddcba")
	println(p)
}
func longestPalindrome2(s string) string {
	if s == "" {
		return  ""
	}
	start,end := 0,0
	for i := 0; i < len(s); i++ {
		left1,right1 := expandAroundCenter(s,i,i)
		left2,right2 := expandAroundCenter(s,i,i + 1)
		if right1 - left1 > end -start {
			start,end = left1,right1
		}
		if right2 - left2 > end -start {
			start,end = left2,right2
		}

	}
	return s[start : end+1]
}
func expandAroundCenter(s string,left,right int)(int,int)  {
	for; left >= 0 && right < len(s) && s[left] == s[right] ;left,right = left - 1,right +1{
	}
	return left+1,right-1
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	left, maxLen := 0, 1
	for i := 0; i < len(s); {
		if len(s)-i <= maxLen/2 {
			break
		}

		begin, end := i, i
		for end < len(s)-1 && s[end+1] == s[end] {
			end++
		}

		i = end + 1
		for end < len(s)-1 && begin > 0 && s[end+1] == s[begin-1] {
			end++
			begin--
		}

		if end+1-begin > maxLen {
			left = begin
			maxLen = end + 1 - begin
		}
	}

	return s[left : left+maxLen]
}
