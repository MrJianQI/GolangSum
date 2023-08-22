//package main
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func swapPairs(head *ListNode) *ListNode {
//	pre := &ListNode{}
//	pre.Next = head
//	behind, one, two := pre, pre, pre
//	for one != nil && two != nil {
//		nxt := two.Next
//
//		behind.Next = two
//		two.Next = one
//		one.Next = nxt
//
//		behind = behind.Next.Next
//		one = nxt
//		if nxt == nil {
//			break
//		}
//		two = nxt.Next
//	}
//	return pre.Next
//}
