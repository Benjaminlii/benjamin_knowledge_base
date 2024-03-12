package main

import (
	"fmt"
	"sync"
)

// 剑指 Offer 12. 矩阵中的路径
// 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。
// 同一个单元格内的字母不允许被重复使用。
// 例如，在下面的 3×4 的矩阵中包含单词 "ABCCED"（单词中的字母已标出）。
// 示例 1：
// 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
// 输出：true

// 示例 2：
// 输入：board = [["a","b"],["c","d"]], word = "abcd"
// 输出：false

func exist(board [][]byte, word string) bool {
	res := false
	wg := sync.WaitGroup{}
	for row, v := range board {
		for col := range v {
			wg.Add(1)
			go func(row, col int) {
				defer wg.Done()
				x := make([][]bool, 0, len(board))
				for i := 0; i < len(board); i++ {
					x = append(x, make([]bool, len(board[i])))
				}
				if dfs(board, row, col, &x, []rune(word)) {
					res = true
				}
			}(row, col)
		}
	}
	wg.Wait()
	return res

}

func dfs(board [][]byte, row, col int, mark *[][]bool, runes []rune) bool {
	if len(runes) == 0 {
		return true
	}
	if row < 0 || row >= len(board) || col < 0 || col >= len(board[row]) {
		return false
	}
	if (*mark)[row][col] {
		return false
	}
	(*mark)[row][col] = true
	defer func() {
		(*mark)[row][col] = false
	}()
	if board[row][col] == byte(runes[0]) {
		return dfs(board, row+1, col, mark, runes[1:]) ||
			dfs(board, row, col+1, mark, runes[1:]) ||
			dfs(board, row-1, col, mark, runes[1:]) ||
			dfs(board, row, col-1, mark, runes[1:])
	} else {
		return false
	}

}

func main() {
	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCCED")) // true
	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCB")) // false
	fmt.Println(exist([][]byte{
		{'a', 'b'},
		{'c', 'd'},
	}, "abcd")) // false
}
