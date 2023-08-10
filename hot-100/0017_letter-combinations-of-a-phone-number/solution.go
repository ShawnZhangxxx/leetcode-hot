/*
__author__ = 'robin-luo'
__date__ = '2023/03/09 10:52'
*/

package main

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
		var temp []string
		for _, y := range res {
			for _, x := range digitMap[string(char)] {
				temp = append(temp, y+string(x))
			}

			res = temp
		}
	}

	return res
}
