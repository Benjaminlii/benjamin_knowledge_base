package main

import "fmt"

// 剑指 Offer 10- I. 斐波那契数列
// 写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）。斐波那契数列的定义如下
// F(0) = 0,   F(1) = 1
// F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
// 斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。
// 答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
func fib(n int) int {
	mod := 1000000007
	if n < 2 {
		return n
	}
	a := 0
	b := 1
	res := 0
	for i := 2; i <= n; i++ {
		res = (a + b) % mod
		a = b
		b = res
	}
	return res
}

func main() {
	fmt.Println(fib(2))
	fmt.Println(fib(5))
}
