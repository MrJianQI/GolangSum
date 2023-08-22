//package main
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	list1, list2 := make([]int, 0), make([]int, 0)
//	a, b := l1, l2
//	for a != nil {
//		list1 = append(list1, a.Val)
//		a = a.Next
//	}
//	for b != nil {
//		list2 = append(list2, b.Val)
//		b = b.Next
//	}
//	m, n := len(list1), len(list2)
//	res := make([]int, 0)
//	i, j := 0, 0
//	temp, flag := 0, 0
//	for i < m && j < n {
//		temp = (list1[i] + list2[j] + flag) % 10
//		res = append(res, temp)
//		flag = (list1[i] + list2[j] + flag) / 10
//		i++
//		j++
//	}
//	for i < m {
//		temp = (list1[i] + flag) % 10
//		res = append(res, temp)
//		flag = (list1[i] + flag) / 10
//		i++
//	}
//	for j < n {
//		temp = (list2[j] + flag) % 10
//		res = append(res, temp)
//		flag = (list2[j] + flag) / 10
//		j++
//	}
//	if flag == 1 {
//		res = append(res, 1)
//	}
//	head := &ListNode{}
//	cur := head
//	for _, v := range res {
//		temp1 := &ListNode{Val: v}
//		cur.Next = temp1
//		cur = cur.Next
//	}
//	return head.Next
//}
