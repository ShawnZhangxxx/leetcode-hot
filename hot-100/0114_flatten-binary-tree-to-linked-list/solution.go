/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-06-28 17:04:57
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-06-28 17:08:30
 * @Description:
 */

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flattern(root *TreeNode) {
	if root == nil {
		return
	}

	flattern(root.Left)
	flattern(root.Right)
	left, right := root.Left, root.Right
	root.Left = nil
	root.Right = left
	cur := root
	for cur.Right != nil {
		cur = cur.Right
	}

	cur.Right = right
}
