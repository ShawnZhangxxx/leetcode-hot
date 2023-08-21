/*
__author__ = 'robin-luo'
__date__ = '2023/03/14 17:10'
*/

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main()  {
	res := groupAnagrams2([]string{"abc","bac","def","fde","cab","a"})
	fmt.Println(res)
}

func groupAnagrams2(strs []string)[][]string  {

	dict := make(map[string][]string)
	for _, str := range strs {
		//将str转化成数组
		strArr := strings.Split(str,"")
		sort.Strings(strArr)
		orderStr := strings.Join(strArr[:],"") //相当于新建切片
		if _, ok := dict[orderStr];ok {
			dict[orderStr] = append(dict[orderStr],str)
		}else {
			dict[orderStr] = []string{str}
		}
	}
	var res  [][]string
	for _, v := range dict {
			res = append(res,v)
	}

	return res
}






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
