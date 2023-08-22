//package main
//
//import "container/heap"
//
//type intHeap []int
//
//func (h intHeap) Len() int           { return len(h) }
//func (h intHeap) Less(i, j int) bool { return h[i] > h[j] }
//func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
//
//func (h *intHeap) Push(val interface{}) {
//	*h = append(*h, val.(int))
//}
//func (h *intHeap) Pop() interface{} {
//	old := *h
//	n := len(old)
//	x := old[n-1]
//	*h = old[0 : n-1]
//	return x
//}
//
//func findKthLargest(nums []int, k int) int {
//	h := &intHeap{}
//	heap.Init(h)
//	for _, i := range nums {
//		heap.Push(h, i)
//	}
//	x := 0
//	for i := 0; i < k; i++ {
//		x = heap.Pop(h).(int)
//	}
//	return x
//}