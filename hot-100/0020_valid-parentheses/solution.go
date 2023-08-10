/*
__author__ = 'robin-luo'
__date__ = '2023/03/09 14:02'
*/

package main

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
