/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-24 11:06:08
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-24 14:00:37
 * @Description:
 */

package main

func removeInvalidParentheses(s string) []string {
	if len(s) == 0 {
		return []string{}
	}

	res := make([]string, 0)
	visited := make(map[string]bool)
	queue := []string{s}
	found := false
	for len(queue) > 0 {
		str := queue[0]
		queue = queue[1:]
		if isValid(str) {
			res = append(res, str)
			found = true
		}

		if found {
			continue
		}

		for i := 0; i < len(str); i++ {
			if str[i] != '(' && str[i] != ')' {
				continue
			}

			nextStr := str[:i] + str[i+1:]
			if !visited[nextStr] {
				visited[nextStr] = true
				queue = append(queue, nextStr)
			}
		}
	}

	return res
}

func isValid(s string) bool {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			count++
		} else if s[i] == ')' {
			count--
			if count < 0 {
				return false
			}
		}
	}

	return count == 0
}
