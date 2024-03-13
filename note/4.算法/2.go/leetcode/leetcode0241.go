package main

import (
	"strconv"
)

// 241. 为运算表达式设计优先级
// 给你一个由数字和运算符组成的字符串 expression ，按不同优先级组合数字和运算符，计算并返回所有可能组合的结果。你可以 按任意顺序 返回答案。
// 生成的测试用例满足其对应输出值符合 32 位整数范围，不同结果的数量不超过 104 。
// 示例 1：
//  输入：expression = "2-1-1"
//  输出：[0,2]
//  解释：
//      ((2-1)-1) = 0
//      (2-(1-1)) = 2
// 示例 2：
//  输入：expression = "2*3-4*5"
//  输出：[-34,-14,-10,-10,10]
//  解释：
//      (2*(3-(4*5))) = -34
//      ((2*3)-(4*5)) = -14
//      ((2*(3-4))*5) = -10
//      (2*((3-4)*5)) = -10
//      (((2*3)-4)*5) = 10
// 提示：
//  1 <= expression.length <= 20
//  expression 由数字和算符 '+'、'-' 和 '*' 组成。
//  输入表达式中的所有整数值在范围 [0, 99]

// Solution 1
func calculate(a int, operator string, b int) int {
	var result int
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	}
	return result
}

func calculateArray(array []string, i int) []int {
	res := []int{}

	if len(array) == 3 {
		a, _ := strconv.Atoi(array[0])
		b, _ := strconv.Atoi(array[2])
		return []int{calculate(a, array[1], b)}
	}

	startIndex := i
	for startIndex+2 < len(array) {
		a, _ := strconv.Atoi(array[startIndex])
		b, _ := strconv.Atoi(array[startIndex+2])
		calculateResult := calculate(a, array[startIndex+1], b)
		nextArray := []string{}
		nextArray = append(nextArray, array[0:startIndex]...)
		nextArray = append(nextArray, strconv.Itoa(calculateResult))
		if startIndex+3 < len(array) {
			nextArray = append(nextArray, array[startIndex+3:]...)
		}

		nextI := startIndex - 2
		if nextI < 0 {
			nextI = 0
		}
		res = append(res, calculateArray(nextArray, nextI)...)
		startIndex += 2
	}

	return res
}

func diffWaysToCompute(expression string) []int {
	input := []string{}
	lp := 0
	for rp := 0; rp < len(expression); rp++ {
		if expression[rp] == '+' || expression[rp] == '-' || expression[rp] == '*' {
			input = append(input, expression[lp:rp])
			input = append(input, string(expression[rp]))
			lp = rp + 1
		}
	}
	input = append(input, expression[lp:])

	if len(input) < 3 {
		res, _ := strconv.Atoi(expression)
		return []int{res}
	}
	return calculateArray(input, 0)
}

// Solution 2
func diffWaysToCompute1(expression string) []int {
	if num, err := strconv.Atoi(expression); err == nil {
		return []int{num}
	}
	ress := []int{}
	for i := 0; i < len(expression); i++ {
		o := expression[i]
		if o != '+' && o != '-' && o != '*' {
			continue
		}
		leftResult := diffWaysToCompute1(expression[:i])
		rightResult := diffWaysToCompute1(expression[i+1:])
		for _, leftNum := range leftResult {
			for _, rightNum := range rightResult {
				res := 0
				switch o {
				case '+':
					res = leftNum + rightNum
				case '-':
					res = leftNum - rightNum
				case '*':
					res = leftNum * rightNum
				}
				ress = append(ress, res)
			}
		}
	}
	return ress
}

func main() {
	res := diffWaysToCompute1("2*3-4*5")
	for _, r := range res {
		println(r)
	}
}
