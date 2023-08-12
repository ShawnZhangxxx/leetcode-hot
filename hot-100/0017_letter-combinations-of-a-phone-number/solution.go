/*
__author__ = 'robin-luo'
__date__ = '2023/03/09 10:52'
*/

package main

import "fmt"

func main() {
	res := letterCombinations2("234")
	fmt.Println(res)
}

func letterCombinations2(digits string) []string {
	digitMap := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	var res []string
	h := map[string]int{}
	for i, v := range digits {

		if i == 0 {
			for _, v2 := range digitMap[string(v)] {
				h[string(v2)] = 1
			}
		} else {
			h2 := map[string]int{}
			for k, _ := range h {
				for _, v2 := range digitMap[string(v)] {
					h2[string(k)+string(v2)] = 1
				}
			}
			h = h2
		}
	}

	for k, _ := range h {
		res = append(res, k)
	}
	return res
}

func letterCombinations(digits string) []string {
	digitMap := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	size := len(digits)
	if size == 0 {
		return []string{}
	}

	res := []string{""}
	for _, char := range digits {
		var temp []string //要设置一个临时数组
		for _, y := range res {
			for _, x := range digitMap[string(char)] {
				temp = append(temp, y+string(x))
			}

			res = temp
		}
	}

	return res
}
