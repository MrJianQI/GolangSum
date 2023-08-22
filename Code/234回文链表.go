//package main
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func isPalindrome(head *ListNode) bool {
//	slow, fast := head, head
//	for fast != nil && fast.Next != nil {
//		slow = slow.Next
//		fast = fast.Next.Next
//	}
//	fast = head
//	temp := reverseList(slow)
//	for temp != nil {
//		if temp.Val != fast.Val {
//			return false
//		}
//		temp = temp.Next
//		fast = fast.Next
//	}
//	return true
//}
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
