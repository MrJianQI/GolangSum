//package main
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	pre := &ListNode{}
//	pre.Next = head
//	slow, fast := pre, pre
//	for i := 0; i <= n; i++ {
//		fast = fast.Next
//	}
//	for fast != nil {
//		fast = fast.Next
//		slow = slow.Next
//	}
//	if slow.Next.Next == nil {
//		slow.Next = nil
//	} else {
//		slow.Next = slow.Next.Next
//	}
//	return slow
//}
