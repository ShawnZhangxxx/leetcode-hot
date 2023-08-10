/*
__author__ = 'robin-luo'
__date__ = '2023/03/14 17:10'
*/

package main

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	dict := make(map[string][]string)
	for _, word := range strs {
		wordSlice := strings.Split(word, "")
		sort.Strings(wordSlice)
		sortedWord := strings.Join(wordSlice[:], "")
		if _, ok := dict[sortedWord]; !ok {
			dict[sortedWord] = []string{word}
		} else {
			dict[sortedWord] = append(dict[sortedWord], word)
		}
	}

	var res [][]string
	for _, v := range dict {
		res = append(res, v)
	}

	return res
}
