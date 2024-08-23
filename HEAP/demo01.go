package main

import (
	"container/heap"
	"fmt"
)

type intHeap []int

// func (sort.Interface) Len() int
// func (sort.Interface) Less(i int, j int) bool
// func (sort.Interface) Swap(i int, j int)

//heap堆常用来实现优先队列


func(h intHeap) Len() int {
	return len(h)
}

func(h *intHeap) Less(i int, j int) bool {
	return (*h)[i]<(*h)[j]
}

func(h *intHeap) Swap(i int, j int)  {
	// var temp int
	temp := (*h)[i]
	(*h)[i] = (*h)[j]
	(*h)[j] = temp
}

func (h *intHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func main()  {
	myIntHeap := &intHeap{8,6,7,2}
	heap.Init(myIntHeap)
	myIntHeap.Push(100)
	len := myIntHeap.Len()
	for i := 0; i < len; i++ {
		fmt.Println(myIntHeap.Pop())
		// fmt.Println("myIntHeap.Pop()")
	}
}