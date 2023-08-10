/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-18 10:38:35
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-18 13:41:25
 * @Description:
 */

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// iteration
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return root
}

// recursion
func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = invertTree1(root.Right), invertTree1(root.Left)
	return root
}
