package main

import "fmt"

// 剑指 Offer 33. 二叉搜索树的后序遍历序列
// 输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。
// 参考以下这颗二叉搜索树：
//	    5
//	   / \
//	  2   6
//	 / \
//	1   3
func verifyPostorder(postorder []int) bool {
	// 空节点和单个节点都视为排序树
	if len(postorder) <= 1 {
		return true
	}

	// 当前数组的最后一个元素是root
	// 向前的连续大于root的应该都为右子树
	root := postorder[len(postorder)-1]
	i := len(postorder) - 2
	for i >= 0 && postorder[i] > root {
		i--
	}

	// 得到左右两颗子树
	leftTree := postorder[:i+1]
	rightTree := postorder[i+1 : len(postorder)-1]

	if len(leftTree) > 0 {
		// 对左子树校验一下
		// 左子树的右半边不能比root大
		leftRoot := postorder[i]
		i--
		for i >= 0 && postorder[i] > leftRoot {
			if postorder[i] > root {
				return false
			}
			i--
		}
	}

	return verifyPostorder(leftTree) && verifyPostorder(rightTree)
}

func main() {
	// 输入: [1,6,3,2,5]
	// 输出: false
	fmt.Println(verifyPostorder([]int{1, 6, 3, 2, 5}))
	// 输入: [1,3,2,6,5]
	// 输出: true
	fmt.Println(verifyPostorder([]int{1, 3, 2, 6, 5}))
	fmt.Println(verifyPostorder([]int{4, 8, 6, 12, 16, 14, 10}))
	fmt.Println(verifyPostorder([]int{1, 2, 5, 10, 6, 9, 4, 3}))
}
