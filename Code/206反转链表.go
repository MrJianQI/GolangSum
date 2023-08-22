//package main
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func reverseList(head *ListNode) *ListNode {
//	if head == nil {
//		return nil
//	}
//	var pre, cur, nxt *ListNode
//	cur = head
//	for cur != nil {
//		nxt = cur.Next
//
//		cur.Next = pre
//		pre = cur
//		cur = nxt
//	}
//	return pre
//
//}
