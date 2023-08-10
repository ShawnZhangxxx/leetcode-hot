/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-27 10:41:21
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-27 10:49:22
 * @Description:
 */

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	return &TreeNode{
		Val:   root1.Val + root2.Val,
		Left:  mergeTrees(root1.Left, root2.Left),
		Right: mergeTrees(root1.Right, root2.Right),
	}
}
