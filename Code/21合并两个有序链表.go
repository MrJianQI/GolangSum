//package main
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//	if list1 == nil {
//		return list2
//	} else if list2 == nil {
//		return list1
//	}
//	head := &ListNode{}
//	cur := head
//	a, b := list1, list2
//	for a != nil && b != nil {
//		if a.Val <= b.Val {
//			temp := &ListNode{
//				Val:  a.Val,
//				Next: nil,
//			}
//			a = a.Next
//			cur.Next = temp
//			cur = cur.Next
//		} else {
//			temp := &ListNode{
//				Val:  b.Val,
//				Next: nil,
//			}
//			b = b.Next
//			cur.Next = temp
//			cur = cur.Next
//		}
//	}
//	if a != nil {
//		cur.Next = a
//	} else if b != nil {
//		cur.Next = b
//	}
//	return head.Next
//}
