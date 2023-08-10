/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-27 09:40:55
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-27 09:53:52
 * @Description:
 */

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	convert(root, &sum)
	return root
}

func convert(node *TreeNode, sum *int) {
	if node == nil {
		return
	}

	convert(node.Right, sum)

	node.Val += *sum
	*sum = node.Val

	convert(node.Left, sum)
}
