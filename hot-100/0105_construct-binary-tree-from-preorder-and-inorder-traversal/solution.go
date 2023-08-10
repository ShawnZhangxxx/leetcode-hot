/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-06-28 15:58:49
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-06-28 16:43:28
 * @Description:
 */

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	preorder = preorder[1:]
	root := &TreeNode{
		Val: rootVal,
	}

	index := 0
	for i, val := range inorder {
		if val == rootVal {
			index = i
			break
		}
	}

	root.Left = buildTree(preorder[:index], inorder[:index])
	root.Right = buildTree(preorder[index:], inorder[index+1:])
	return root
}
