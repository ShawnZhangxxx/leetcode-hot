/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-10 10:41:27
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-10 10:58:23
 * @Description:
 */

package main

import "fmt"

func main() {
	res := canFinish2(4, [][]int{
		{3, 1},
		{2, 1},
		{3, 2},
		{1, 3},
	})
	fmt.Println(res)
}
func canFinish2(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	queue := make([]int, 0)
	inDegree := make([]int, numCourses)

	for i := 0; i < len(prerequisites); i++ {
		src := prerequisites[i][1]
		dst := prerequisites[i][0]
		if graph[src] == nil {
			graph[src] = []int{}
		}
		graph[src] = append(graph[src], dst)
		inDegree[dst] ++
	}

	//第一次初始化队列
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		peek := queue[0]
		queue = queue[1:]
		for j := 0; j < len(graph[peek]); j++ {
			inDegree[graph[peek][j]] --
			if inDegree[graph[peek][j]] == 0 {
				queue = append(queue, graph[peek][j])
			}
		}
	}
	for i := 0; i < len(inDegree); i++ {
		if inDegree[i] > 0 {
			return  false
		}
	}
	return true
}

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
