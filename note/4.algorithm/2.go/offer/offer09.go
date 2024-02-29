package offer

// 剑指 Offer 09. 用两个栈实现队列
// 用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，
// 分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

type CQueue struct {
    InputStack  []int
    OutputStack []int
}

func Constructor() CQueue {
    return CQueue{}
}

func (cq *CQueue) AppendTail(value int) {
    cq.InputStack = append(cq.InputStack, value)
}

func (cq *CQueue) DeleteHead() int {
    if len(cq.OutputStack) == 0 {
        for i := len(cq.InputStack) - 1; i >= 0; i-- {
            cq.OutputStack = append(cq.OutputStack, cq.InputStack[i])
        }
        cq.InputStack = make([]int, 0)
    }
    if len(cq.OutputStack) == 0 {
        return -1
    }
    res := cq.OutputStack[len(cq.OutputStack)-1]
    cq.OutputStack = cq.OutputStack[:len(cq.OutputStack)-1]
    return res
}
