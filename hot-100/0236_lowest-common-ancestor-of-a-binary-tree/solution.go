/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-18 14:02:05
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-18 14:04:30
 * @Description:
 */

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l != nil && r != nil {
		return root
	} else if l != nil {
		return l
	} else {
		return r
	}
}
