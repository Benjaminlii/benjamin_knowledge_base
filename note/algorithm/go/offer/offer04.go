package main

import "fmt"

// 剑指 Offer 04. 二维数组中的查找
// 在一个 n * m 的二维数组中，每一行都按照从左到右 非递减 的顺序排序，每一列都按照从上到下 非递减 的顺序排序。
// 请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
// 示例:
// 现有矩阵 matrix 如下：
// [
//
//	[1,   4,  7, 11, 15],
//	[2,   5,  8, 12, 19],
//	[3,   6,  9, 16, 22],
//	[10, 13, 14, 17, 24],
//	[18, 21, 23, 26, 30]
//
// ]
// 给定 target = 5，返回 true。
// 给定 target = 20，返回 false。
func findNumberIn2DArray(matrix [][]int, target int) bool {
	res := false
	for i, j := len(matrix)-1, 0; i >= 0 && j < len(matrix[0]); {
		getNum := matrix[i][j]
		if getNum == target {
			res = true
			break
		} else if getNum > target {
			i--
		} else {
			j++
		}
	}
	return res
}

func main() {
	fmt.Println(findNumberIn2DArray([][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}, 7))
	fmt.Println(findNumberIn2DArray([][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}, 3))
}
