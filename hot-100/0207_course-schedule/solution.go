/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-10 10:41:27
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-10 10:58:23
 * @Description:
 */

package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	list := make([][]int, numCourses)
	inDegree := make([]int, numCourses)
	queue := make([]int, 0)

	for i := range prerequisites {
		src, dst := prerequisites[i][0], prerequisites[i][1]
		inDegree[dst]++
		if list[src] == nil {
			list[src] = []int{}
		}

		list[src] = append(list[src], dst)
	}

	for i, v := range inDegree {
		if v == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) != 0 {
		c := queue[0]
		queue = queue[1:]
		for _, dst := range list[c] {
			inDegree[dst]--
			if inDegree[dst] == 0 {
				queue = append(queue, dst)
			}
		}

		numCourses--
	}

	return numCourses == 0
}
