package main

import "fmt"

// 剑指 Offer 05. 替换空格
// 请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
// 示例 1：
// 输入：s = "We are happy."
// 输出："We%20are%20happy."
func replaceSpace(s string) string {
	res := []rune{}
	for _, r := range s {
		if r == rune(' ') {
			res = append(res, []rune("%20")...)
		} else {
			res = append(res, r)
		}
	}
	return string(res)
}

func main() {
	fmt.Println(replaceSpace("We Are Happy"))
	fmt.Println(replaceSpace("1 二 three xxx"))
}
