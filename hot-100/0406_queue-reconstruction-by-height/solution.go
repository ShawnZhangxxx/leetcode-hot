/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-26 09:39:30
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-26 10:12:51
 * @Description:
 */

package main

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}

		return people[i][0] > people[j][0]
	})

	res := make([][]int, 0)
	for _, p := range people {
		if p[1] >= len(res) {
			res = append(res, p)
		} else {
			res = append(res[:p[1]], append([][]int{p}, res[p[1]:]...)...)
		}
	}

	return res
}
