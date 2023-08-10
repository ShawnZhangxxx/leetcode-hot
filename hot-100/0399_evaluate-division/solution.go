/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-25 14:55:01
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-25 15:22:21
 * @Description:
 */

package main

type Node struct {
	dest   string
	weight float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string][]Node)
	for i, equation := range equations {
		a, b := equation[0], equation[1]
		value := values[i] // a/b
		graph[a] = append(graph[a], Node{dest: b, weight: value})
		graph[b] = append(graph[b], Node{dest: a, weight: 1 / value})
	}

	res := make([]float64, len(queries))
	for i, query := range queries {
		start, end := query[0], query[1]
		_, ok1 := graph[start]
		_, ok2 := graph[end]
		if !ok1 || !ok2 {
			res[i] = -1.0 // 若开始或结束节点不存在，则无法计算结果
		} else {
			visited := make(map[string]bool)
			res[i] = dfs(graph, start, end, 1.0, visited)
		}
	}

	return res
}

func dfs(graph map[string][]Node, curNode, endNode string, curRes float64, visited map[string]bool) float64 {
	if curNode == endNode {
		return curRes
	}

	visited[curNode] = true
	for _, node := range graph[curNode] {
		if !visited[node.dest] {
			res := dfs(graph, node.dest, endNode, curRes*node.weight, visited)
			if res != -1 {
				return res
			}
		}
	}

	return -1.0
}
