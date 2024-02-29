package main

import (
    "fmt"
    "math/rand"
    "time"
)

// 快速排序
func QuickSort(array []int) {
    defer fmt.Println(time.Now())
    quickSort(array, 0, len(array)-1)

}

func quickSort(array []int, startIndex, endIndex int) {
    if startIndex < endIndex {
        index := partition(array, startIndex, endIndex)
        go quickSort(array, startIndex, index-1)
        go quickSort(array, index+1, endIndex)
    }
}

func partition(array []int, startIndex, endIndex int) int {
    mark := array[startIndex]
    for startIndex < endIndex {
        for array[endIndex] >= mark && startIndex < endIndex {
            endIndex--
        }
        array[startIndex] = array[endIndex]
        for array[startIndex] <= mark && startIndex < endIndex {
            startIndex++
        }
        array[endIndex] = array[startIndex]
    }
    array[startIndex] = mark
    return startIndex
}

func createSortReq(length, start, end int) []int {
    res := make([]int, 0, length)

    rand.Seed(time.Now().Unix())
    for i := 0; i < length; i++ {
        random := rand.Int()
        random %= end - start
        random += start
        res = append(res, random)
    }

    return res
}
func main() {
    createSortReq := func(length, start, end int) []int {
        res := make([]int, 0, length)
        rand.Seed(time.Now().Unix())
        for i := 0; i < length; i++ {
            random := rand.Int()
            random %= end - start
            random += start
            res = append(res, random)
        }
        return res
    }
    req := createSortReq(500000, 0, 500000)
    // fmt.Println(req)
    QuickSort(req)
    // fmt.Println(req)
}
