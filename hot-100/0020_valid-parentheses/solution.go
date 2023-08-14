/*
__author__ = 'robin-luo'
__date__ = '2023/03/09 14:02'
*/

package main

import "fmt"

func main() {
	fmt.Println(isValid2("({[]})[]"))
}

func isValid2(s string) bool {
	var stack []byte
	h := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, byte(v))
		} else {
			pop := stack[len(stack)-1]
			if pop != h[byte(v)] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0

}

func isValid(s string) bool {
	var stack []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case '(', '[', '{':
			stack = append(stack, c)
		case ')', ']', '}':
			size := len(stack)
			if size == 0 {
				return false
			}

			v := stack[size-1]
			stack = stack[:size-1]
			if c == ')' && v != '(' {
				return false
			}

			if c == ']' && v != '[' {
				return false
			}

			if c == '}' && v != '{' {
				return false
			}
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}
