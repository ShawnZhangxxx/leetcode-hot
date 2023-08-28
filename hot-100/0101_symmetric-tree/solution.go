/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-06-28 14:21:24
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-06-28 14:43:27
 * @Description:
 */

package main

import (
	"fmt"
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
	root = initTreeLayer(10,8,8,7,7,7,7)
	printTreeLayer(root)
	res := isSymmetric2(root)
	fmt.Println(res)
}
//初始化一个搜索二叉树
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
func initTreeLayer( Data ...int) *TreeNode {
	var queue []*TreeNode
	head := &TreeNode{
		Val: Data[0],
	}
	queue =  append(queue,head)
	i := 0
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		i = i+1
		if i<len(Data) {
			top.Left = &TreeNode{Val: Data[i]}
			queue = append(queue,top.Left)
		}else {
			break
		}
		i = i+1
		if i<len(Data) {
			top.Right = &TreeNode{Val: Data[i]}
			queue = append(queue,top.Right)
		}else {
			break
		}
	}
	return  head
}
//层序遍历
func printTreeLayer( root *TreeNode)  {
	var queue []*TreeNode
	queue =  append(queue,root)
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		fmt.Println(top.Val)
		if top.Left != nil {
			queue = append(queue,top.Left)
		}
		if top.Right != nil {
			queue = append(queue,top.Right)
		}
	}
}
//迭代校验对称
func isSymmetric3(root *TreeNode) bool {
	u, v := root, root
	q := []*TreeNode{}
	q = append(q, u)
	q = append(q, v)
	for len(q) > 0 {
		u, v = q[0], q[1]
		q = q[2:]
		if u == nil && v == nil { continue }
		if u == nil || v == nil { return false }
		if u.Val != v.Val { return false }
		q = append(q, u.Left)
		q = append(q, v.Right)
		q = append(q, u.Right)
		q = append(q, v.Left) }
	return true
}



func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return helper2(root.Left, root.Right)
}
func helper2(left *TreeNode,right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil  {
			return  false
	}
	if left.Val != right.Val {
		return  false
	}
	return helper2(left.Left,right.Right) && helper(left.Right,right.Left)
}








func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return helper(root.Left, root.Right)
}

func helper(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return helper(left.Left, right.Right) && helper(left.Right, right.Left)
}
