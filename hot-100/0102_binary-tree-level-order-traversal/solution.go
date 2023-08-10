/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-06-28 14:45:33
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-06-28 15:10:17
 * @Description:
 */

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// dfs
func levelOrder(root *TreeNode) [][]int {
	var (
		res [][]int
		dfs func(root *TreeNode, level int)
	)

	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if level >= len(res) {
			res = append(res, []int{})
		}

		res[level] = append(res[level], root.Val)
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)
	return res
}

// 迭代
func levelOrder1(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		tempArr := make([]int, len(q))
		for i := range tempArr {
			node := q[0]
			q = q[1:]
			tempArr[i] = node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		res = append(res, tempArr)
	}

	return res
}
