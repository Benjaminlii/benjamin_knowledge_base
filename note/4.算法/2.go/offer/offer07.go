package offer

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 剑指 Offer 07. 重建二叉树
// 输入某二叉树的前序遍历和中序遍历的结果，请构建该二叉树并返回其根节点。
// 假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
// 示例 1:
// Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
// Output: [3,9,20,null,null,15,7]
// 示例 2:
// Input: preorder = [-1], inorder = [-1]
// Output: [-1]
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder)+len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}

	var leftVin []int
	var rightVin []int
	var leftPre []int
	var rightPre []int
	for i, v := range inorder {
		if v == root.Val {
			leftVin = inorder[0:i]
			rightVin = inorder[i+1:]
			leftPre = preorder[1 : len(leftVin)+1]
			rightPre = preorder[1+len(leftPre):]
			break
		}
	}
	root.Left = buildTree(leftPre, leftVin)
	root.Right = buildTree(rightPre, rightVin)
	return root
}

func main() {
	x := buildTree([]int{1, 2, 4, 7, 3, 5, 6, 8}, []int{4, 7, 2, 1, 5, 3, 8, 6})
	fmt.Println(x)
}
