//package main
//
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//func inorderTraversal(root *TreeNode) []int {
//	res := make([]int, 0)
//	cur := root
//	stack := make([]*TreeNode, 0)
//
//	for len(stack) != 0 || cur != nil {
//		if cur != nil {
//			stack = append(stack, cur)
//			cur = cur.Left
//			continue
//		}
//		cur = stack[len(stack)-1]
//		res = append(res, cur.Val)
//		cur = cur.Right
//		stack = stack[:len(stack)-1]
//
//	}
//	return res
//}
