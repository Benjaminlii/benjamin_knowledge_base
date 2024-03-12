package main

// 199. 二叉树的右视图
// 给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
//	   5
//	  / \
//	 2   6
//	/ \
// 1   3
// output: [5,6,3]
// 示例 1:
//  输入: [1,2,3,null,5,null,4]
//  输出: [1,3,4]
// 示例 2:
//  输入: [1,null,3]
//  输出: [1,3]
// 示例 3:
//  输入: []
//  输出: []

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	// TODO 待优化
	if root == nil {
		return nil
	}
	leftResult := rightSideView(root.Left)
	rightResult := rightSideView(root.Right)
	if len(leftResult) > len(rightResult){
		rightResult = append(rightResult, leftResult[len(rightResult):]...)
	}
	return append([]int{root.Val}, rightResult...)
}

func main() {
	root := &TreeNode{
		Val:   5,
		Left:  &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   1,
			},
			Right: &TreeNode{
				Val:   3,
			},
		},
		Right: &TreeNode{
			Val:   6,
		},
	}
	res := rightSideView(root)
	for _, v := range res {
		println(v)
	}
}
