/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-06-28 10:47:32
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-06-28 10:58:11
 * @Description:
 */

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//       5
//   3	    	7
// 2   4
func main() {

	var root *TreeNode
	root = insert(root, 5)
	root = insert(root, 3)
	root = insert(root, 7)
	root = insert(root, 4)
	root = insert(root, 2)

	// res := inorderTraversal2(root)
	res := inorderTraversal3(root)
	fmt.Println(res)
}

//初始化一个二叉树  层序
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

func inorderTraversal2(root *TreeNode) []int { //递归 中序遍历是顺序遍历
	var res []int
	if root == nil {
		return nil
	}
	res = append(res, inorderTraversal2(root.Left)...)
	fmt.Println(root.Val)
	res = append(res, root.Val)
	res = append(res, inorderTraversal2(root.Right)...)
	return res

}

func inorderTraversal3(root *TreeNode) []int { //迭代
	var (
		stack []*TreeNode
		res   []int
	)
	// stack = append(stack, root)
	// for len(stack) > 0 {
	// 	for root != nil && root.Left != nil {
	// 		stack = append(stack, root.Left)
	// 		root = root.Left
	// 	}
	// 	pop := stack[len(stack)-1]
	// 	stack = stack[:len(stack)-1]
	// 	res = append(res, pop.Val)
	// 	root = pop.Right
	// 	if pop.Right != nil {
	// 		stack = append(stack, pop.Right)
	// 	}
	// }

	for {
		for root != nil { //比自己的方法的区别是，判断root的值，而不是root.left，节俭的很多代码，而且包含初始化条件
			stack = append(stack, root)
			root = root.Left
		}

		if len(stack) == 0 {
			return res
		}
		pop := stack[len(stack)-1]
		res = append(res, pop.Val)
		fmt.Println(pop.Val)
		stack = stack[:len(stack)-1]
		root = pop.Right
	}
}

// 迭代
func inorderTraversal(root *TreeNode) []int {
	var (
		stack []*TreeNode
		res   []int
	)

	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		size := len(stack)
		if size == 0 {
			return res
		}

		node := stack[size-1]
		stack = stack[:size-1]
		res = append(res, node.Val)
		root = node.Right
	}
}

// 递归
func inorderTraversal1(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	res = append(res, inorderTraversal1(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal1(root.Right)...)
	return res
}
