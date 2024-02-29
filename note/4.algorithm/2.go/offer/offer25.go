package main

import "fmt"

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

// 剑指 Offer 25. 合并两个排序的链表
// 输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
// 输入：1->2->4, 1->3->4
// 输出：1->1->2->3->4->4
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var res, h11, h12, h2 *ListNode
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    if l1.Val <= l2.Val {
        h11 = l1
        h12 = l1.Next
        h2 = l2
    } else {
        h11 = l2
        h12 = l2.Next
        h2 = l1
    }
    res = h11
    for h12 != nil && h2 != nil {
        if h2.Val >= h11.Val && h2.Val < h12.Val {
            m := h2
            h2 = h2.Next

            m.Next = h12
            h11.Next = m
            h11 = m
            continue
        }
        h11 = h12
        h12 = h12.Next
    }
    if h2 != nil {
        if h12 != nil {
            h12.Next = h2
        } else {
            h11.Next = h2
        }
    }
    return res
}

func main() {
    // 输入：1->2->4, 1->3->4
    // 输出：1->1->2->3->4->4
    fmt.Println(mergeTwoLists(&ListNode{
        Val: 1,
        Next: &ListNode{
            Val: 2,
            Next: &ListNode{
                Val: 4,
            },
        },
    }, &ListNode{
        Val: 1,
        Next: &ListNode{
            Val: 3,
            Next: &ListNode{
                Val: 4,
            },
        },
    }))
}
