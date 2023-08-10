/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-25 14:21:05
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-25 14:53:41
 * @Description:
 */

package main

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	stack := []string{}
	curNum := 0
	curStr := ""

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num, _ := strconv.Atoi(string(s[i]))
			curNum = curNum*10 + num
		} else if s[i] == '[' {
			stack = append(stack, curStr)
			stack = append(stack, strconv.Itoa(curNum))
			curNum = 0
			curStr = ""
		} else if s[i] == ']' {
			num, _ := strconv.Atoi(stack[len(stack)-1])
			prevStr := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			curStr = prevStr + strings.Repeat(curStr, num)
		} else {
			curStr += string(s[i])
		}
	}

	return curStr
}
