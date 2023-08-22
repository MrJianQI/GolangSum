package main

import (
	"container/heap"
	"fmt"
)

type Meta struct {
	Val   int
	Count int
}
type MaxList []*Meta

func (h MaxList) Len() int {
	return len(h)
}
func (h MaxList) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	return
}
func (h MaxList) Less(i, j int) bool {
	return h[i].Count > h[j].Count
}
func (h *MaxList) Push(x interface{}) {
	(*h) = append(*h, x.(*Meta))
	return
}
func (h *MaxList) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	list := make(MaxList, 0)
	heap.Init(&list)
	dict := make(map[int]int)
	for _, v := range nums {
		if _, ok := dict[v]; !ok {
			dict[v] = 1
		} else {
			dict[v] += 1
		}
	}
	for i, v := range dict {
		heap.Push(&list, &Meta{Val: i, Count: v})
	}
	res := make([]int, 0)
	for i := 0; i < k; i++ {
		temp := heap.Pop(&list).(*Meta)
		res = append(res, temp.Val)
	}
	return res
}
func main() {
	res := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	fmt.Println(res)

}
