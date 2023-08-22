//package main
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func getIntersectionNode(headA, headB *ListNode) *ListNode {
//	if headA == nil || headB == nil {
//		return nil
//	}
//	l, r := headA, headB
//	for l != nil || r != nil {
//		if l == r {
//			break
//		}
//		if l == nil {
//			l = headB
//		} else {
//			l = l.Next
//		}
//		if r == nil {
//			r = headA
//		} else {
//			r = r.Next
//		}
//	}
//	return l
//}
