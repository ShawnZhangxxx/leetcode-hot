/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-06-28 15:58:49
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-06-28 16:43:28
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
	//       3
	//   9	    	20
	// `         15     7
	preorder := []int{3,9,20,15,7}
	inorder := []int{9,3,15,20,7}
	res := buildTree2(preorder,inorder)
	printTreeLayer(res)
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

//先序   root 左 右
//中序   左 root 右
func buildTree2(preorder,inorder []int) *TreeNode{
	if len(preorder) == 0 {
		return 	nil
	}
	var index  = 0
	for i := 0; i < len(inorder); i++ {
		if preorder[0] == inorder[i] {
			index = i
			break
		}
	}
	root := &TreeNode{
		Val: preorder[0],
	}					//这个因为 先序和中序的长度一样长,所以也能决定先序的长度
	root.Left = buildTree2(preorder[1:index+1],inorder[:index])
	root.Right = buildTree2(preorder[index+1:],inorder[index+1:])
	return root
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
