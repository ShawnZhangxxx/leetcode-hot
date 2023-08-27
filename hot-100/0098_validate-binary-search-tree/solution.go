/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-06-28 14:00:44
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-06-28 14:38:43
 * @Description:
 */

package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//       5
	//   3	    	7
	// 2   4
	var root *TreeNode
	root = insert(root, 5)
	root = insert(root, 3)
	root = insert(root, 7)
	root = insert(root, 4)
	root = insert(root, 2)
	root.Left.Left.Left = &TreeNode{
		Val: 8,
	}
	res := isValidBST2(root)
	fmt.Println(res)

}

//初始化一个二叉树
func insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = insert(root.Left, val)
	} else {
		root.Right = insert(root.Right, val)
	}
	return root

}

func isValidBST2(root *TreeNode) bool {
	return isValid(root)
}

func isValid(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Left.Val > root.Val {
		return false
	} else if root.Right != nil && root.Right.Val < root.Val {
		return false
	}
	return isValid(root.Left) && isValid(root.Right)
}

func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(root *TreeNode, left, right int) bool {
	if root == nil {
		return true
	}

	if left >= root.Val || root.Val >= right {
		return false
	}

	return validate(root.Left, left, root.Val) && validate(root.Right, root.Val, right)
}
