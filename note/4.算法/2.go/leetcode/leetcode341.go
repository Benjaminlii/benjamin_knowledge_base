package main

// NestedInteger [[1,2],2,[1,[1,2,3]]]
type NestedInteger struct {
    value  int
    isList bool

    nesteds []*NestedInteger
}
func (this *NestedInteger) IsInteger() bool {
    return !this.isList
}
func (this *NestedInteger) GetInteger() int {
    if this.isList {
        panic("NestedInteger holds a nested list, cannot get integer")
    }
    return this.value
}
func (this *NestedInteger) SetInteger(value int) {
    this.value = int(value)
    this.isList = false
    this.nesteds = nil
}
func (this *NestedInteger) Add(elem NestedInteger) {
    if this.isList {
        this.nesteds = append(this.nesteds, &elem)
    } else {
        this.isList = true
        this.nesteds = []*NestedInteger{&elem}
    }
}
func (this *NestedInteger) GetList() []*NestedInteger {
    return this.nesteds
}

// 给你一个嵌套的整数列表 nestedList 。每个元素要么是一个整数，要么是一个列表；该列表的元素也可能是整数或者是其他列表。
// 请你实现一个迭代器将其扁平化，使之能够遍历这个列表中的所有整数。
// 实现扁平迭代器类 NestedIterator ：
// NestedIterator(List<NestedInteger> nestedList) 用嵌套列表 nestedList 初始化迭代器。
// int next() 返回嵌套列表的下一个整数。
// boolean hasNext() 如果仍然存在待迭代的整数，返回 true ；否则，返回 false 。
// 你的代码将会用下述伪代码检测：
// initialize iterator with nestedList
// res = []
// while iterator.hasNext()
//     append iterator.next() to the end of res
// return res
// 如果 res 与预期的扁平化列表匹配，那么你的代码将会被判为正确。
// 示例 1：
// 输入：nestedList = [[1,1],2,[1,1]]
// 输出：[1,1,2,1,1]
// 解释：通过重复调用 next 直到 hasNext 返回 false，next 返回的元素的顺序应该是: [1,1,2,1,1]。
// 
// 示例 2：
// 输入：nestedList = [1,[4,[6]]]
// 输出：[1,4,6]
// 解释：通过重复调用 next 直到 hasNext 返回 false，next 返回的元素的顺序应该是: [1,4,6]。

type NestedIterator struct {
    // 将列表视作一个队列，栈中直接存储该队列
    stack [][]*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    return &NestedIterator{[][]*NestedInteger{nestedList}}
}

func (it *NestedIterator) Next() int {
    // 由于保证调用 Next 之前会调用 HasNext，直接返回栈顶列表的队首元素，将其弹出队首并返回
    queue := it.stack[len(it.stack)-1]
    val := queue[0].GetInteger()
    it.stack[len(it.stack)-1] = queue[1:]
    return val
}

func (it *NestedIterator) HasNext() bool {
    for len(it.stack) > 0 {
        queue := it.stack[len(it.stack)-1]
        if len(queue) == 0 { // 当前队列为空，出栈
            it.stack = it.stack[:len(it.stack)-1]
            continue
        }
        nest := queue[0]
        if nest.IsInteger() {
            return true
        }
        // 若队首元素为列表，则将其弹出队列并入栈
        it.stack[len(it.stack)-1] = queue[1:]
        it.stack = append(it.stack, nest.GetList())
    }
    return false
}



func main() {
	input := []*NestedInteger{
		{value: 1},
		{value: 2},
		{isList: true, nesteds: []*NestedInteger{
			{value: 3},
			{value: 4},
		}},
		{isList: true, nesteds: []*NestedInteger{
			{value: 5},
			{isList: true, nesteds: []*NestedInteger{
				{value: 6},
				{value: 7},
			}},
		}},
	}
	iterator := Constructor(input)
	res := []int{}
	for iterator.HasNext() {
		res = append(res, iterator.Next())
	}

    for _, r := range res {
        print(r)
    }
}
