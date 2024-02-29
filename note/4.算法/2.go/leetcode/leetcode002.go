package main

import "fmt"

// leetcode 2 两数相加
// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
// 示例 1：
// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807.

// 示例 2：
// 输入：l1 = [0], l2 = [0]
// 输出：[0]

// 示例 3：
// 输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// 输出：[8,9,9,9,0,0,0,1]

type ListNode struct {
    Val  int
    Next *ListNode
}

func (head *ListNode) String() string {
    nextStr := "nil"
    if head.Next != nil {
        nextStr = head.Next.String()
    }
    return fmt.Sprintf("%d->%s", head.Val, nextStr)
}

// 基本思路，顺序遍历，合并两链表
func addTwoNumbers0(l1 *ListNode, l2 *ListNode) *ListNode {
    ans := &ListNode{}
    p := ans
    add := 0
    for (l1 != nil && l2 != nil) || add != 0 {
        val := 0
        if l1 != nil {
            val += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            val += l2.Val
            l2 = l2.Next
        }
        val += add
        add = 0
        if val >= 10 {
            val -= 10
            add = 1
        }
        p.Next = &ListNode{Val: val}
        p = p.Next
    }
    if l1 != nil {
        p.Next = l1
    }
    if l2 != nil {
        p.Next = l2
    }
    return ans.Next
}

// 递归思路
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    } else if l2 == nil {
        return l1
    }

    sum := l1.Val + l2.Val
    if sum >= 10 {
        return &ListNode{
            Val: sum - 10,
            Next: addTwoNumbers1(
                &ListNode{Val: 1},
                addTwoNumbers0(l1.Next, l2.Next),
            )}
    } else {
        return &ListNode{
            Val:  sum,
            Next: addTwoNumbers1(l1.Next, l2.Next),
        }
    }
}

func main() {
    l1 := &ListNode{Val: 9}
    l1.Next = &ListNode{Val: 9}
    l1.Next.Next = &ListNode{Val: 9}
    l1.Next.Next.Next = &ListNode{Val: 9}
    l1.Next.Next.Next.Next = &ListNode{Val: 9}
    l1.Next.Next.Next.Next.Next = &ListNode{Val: 9}
    l1.Next.Next.Next.Next.Next.Next = &ListNode{Val: 9}
    l2 := &ListNode{Val: 9}
    l2.Next = &ListNode{Val: 9}
    l2.Next.Next = &ListNode{Val: 9}
    l2.Next.Next.Next = &ListNode{Val: 9}

    fmt.Println(addTwoNumbers1(l1, l2))
}
