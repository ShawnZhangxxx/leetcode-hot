/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-26 10:33:21
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-26 10:41:51
 * @Description:
 */

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	res := 0
	if root == nil {
		return res
	}

	res = dfs(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}

func dfs(root *TreeNode, targetSum int) int {
	res := 0
	if root == nil {
		return res
	}

	if root.Val == targetSum {
		res++
	}

	res += dfs(root.Left, targetSum-root.Val)
	res += dfs(root.Right, targetSum-root.Val)
	return res
}
