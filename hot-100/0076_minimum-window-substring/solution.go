/*
__author__ = 'robin-luo'
__date__ = '2023/06/26 15:03'
*/

package main

import "fmt"

func main()  {
	res := minWindow("apsa","as")
	fmt.Println(res)
}
func minWindow2(s, t string) string {
	if len(s) < len(t) {
		return ""
	}
	needNum := len(t)
	hash := map[string]int{}
	for i := 0; i < len(t); i++ {
		hash[string(t[i])] = hash[string(t[i])]+1
	}
	l,r := 0,0
	lmin := 0
	rmin := len(s)
	for r < len(s) {
		if _,ok := hash[string(s[r])] ;ok  {//命中
			if hash[string(s[r])] > 0{
				needNum--
			}
			hash[string(s[r])] = hash[string(s[r])] -1
		}
		//收缩
		if needNum == 0 {
			for i := l; i <= r; i++ {
				if v,ok := hash[string(s[i])] ;ok  { //命中
					if v ==  0 {//正好收缩到了
						if r - i <  rmin - lmin {
							lmin = i
							rmin = r
						}
						l = i+1
						needNum ++
						hash[string(s[i])] ++
						break
					}else {
						hash[string(s[i])] ++
					}
				}
			}
		}
		r++
	}
	if rmin == len(s) {
		return ""
	}else {
		return s[lmin:rmin+1]
	}
}

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

		for l < i && m[s[l]] < 0 { //只有 扣成-的才会触发 这个for
			m[s[l]]++
			l++
		}

		if count == 0 && max > i-l+1 { //达到覆盖全部切 最大值长度<max
			max = i - l + 1
			res = s[l : i+1]
		}
	}

	return res
}
