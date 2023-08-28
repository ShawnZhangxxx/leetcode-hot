/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-06-28 14:45:33
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-06-28 15:10:17
 * @Description:
 */

package main

import "fmt"

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
	root = initTreeLayer(10,8,8,7,7,7,7,0)
	printTreeLayer(root)
	res := levelOrder3(root)
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

//层序遍历
func levelOrder3(root *TreeNode) [][]int {
	var curQueue []*TreeNode
	var NextQueue []*TreeNode
	var res [][]int
	var curLevel []int
	curQueue = append(curQueue,root)
	for len(curQueue) > 0 {
		top := curQueue[0]
		curQueue = curQueue[1:]
		curLevel = append(curLevel,top.Val)

		if top.Left != nil {
			NextQueue = append(NextQueue,top.Left)
		}
		if top.Right != nil {
			NextQueue = append(NextQueue,top.Right)
		}
		if len(curQueue) == 0  {
			res = append(res,curLevel)
			curLevel = nil
			if len(NextQueue) > 0 {
				curQueue = NextQueue
				NextQueue = nil
			}
		}
	}
	return  res
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
		if level >= len(res) { //这段代码 防止golang数组溢出的初始化代码.
			res = append(res, []int{})
		}
		res[level] = append(res[level], root.Val)
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)
	return res
}

// 迭代 这个方法利用了tempArr数量正好是每一层的数量,简化了很多代码,多使用一个每层的局部变量
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
