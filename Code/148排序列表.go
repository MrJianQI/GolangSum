package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	fast = slow.Next
	slow.Next = nil
	slow = head
	slow = sortList(slow)
	fast = sortList(fast)
	cur := merge(slow, fast)
	return cur
}
func merge(list1, list2 *ListNode) *ListNode {
	if list2 == nil {
		return list1
	} else if list1 == nil {
		return list1
	} else {
		head := &ListNode{}
		cur := head
		for list1 != nil && list2 != nil {
			if list1.Val <= list2.Val {
				temp := &ListNode{Val: list1.Val}
				cur.Next = temp
				cur = cur.Next
				list1 = list1.Next
			} else {
				temp := &ListNode{Val: list2.Val}
				cur.Next = temp
				cur = cur.Next
				list2 = list2.Next
			}
		}
		if list1 != nil {
			cur.Next = list1
		}
		if list2 != nil {
			cur.Next = list2
		}
		return head.Next
	}
}
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{Val: 0, Next: nil}
	tail := head
	for list1 != nil && list2 != nil {
		if list1.Val == list2.Val {
			t1, t2 := list1.Next, list2.Next
			list1.Next, list2.Next = list2, nil
			tail.Next = list1
			tail = tail.Next.Next
			list1, list2 = t1, t2
		} else if list1.Val < list2.Val {
			t1 := list1.Next
			list1.Next = nil
			tail.Next = list1
			tail = tail.Next
			list1 = t1
		} else {
			t1 := list2.Next
			list2.Next = nil
			tail.Next = list2
			tail = tail.Next
			list2 = t1
		}
	}
	if list1 == nil && list2 != nil {
		tail.Next = list2
	}
	if list1 != nil && list2 == nil {
		tail.Next = list1
	}
	return head.Next
}
func main() {
	head := &ListNode{Val: 4}
	cur1 := &ListNode{Val: 2}
	cur2 := &ListNode{Val: 1}
	cur3 := &ListNode{Val: 3}
	head.Next = cur1
	//cur1.Next = cur2
	cur2.Next = cur3
	//cur := sortList(head)
	cur := mergeTwoLists(head, cur2)
	for cur != nil {
		fmt.Print(cur.Val)
		fmt.Print(" ")
		cur = cur.Next
	}
}
