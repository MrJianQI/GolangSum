//package main
//
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//func levelOrder(root *TreeNode) [][]int {
//	if root == nil {
//		return [][]int{}
//	} else {
//		res := make([][]int, 0)
//		stack := make([]*TreeNode, 0)
//		cur := root
//		stack = append(stack, cur)
//		for len(stack) != 0 {
//			length := len(stack)
//			temp := make([]int, 0)
//			for i := 0; i < length; i++ {
//				if stack[i] == nil {
//					continue
//				}
//				temp = append(temp, stack[i].Val)
//				if stack[i].Left != nil {
//					stack = append(stack, stack[i].Left)
//				}
//				if stack[i].Right != nil {
//					stack = append(stack, stack[i].Right)
//				}
//			}
//			res = append(res, temp)
//			stack = stack[length:]
//		}
//		return res
//	}
//}
