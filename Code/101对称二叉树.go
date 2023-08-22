//package main
//
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//func isSymmetric(root *TreeNode) bool {
//	if root == nil {
//		return true
//	} else if root.Left != nil && root.Right == nil {
//		return false
//	} else if root.Left == nil && root.Right != nil {
//		return false
//	} else if root.Left == nil && root.Right == nil {
//		return true
//	} else {
//		if root.Left.Val != root.Right.Val {
//			return false
//		}
//		return isOk(root.Left, root.Right)
//	}
//}
//func isOk(a, b *TreeNode) bool {
//	if a == nil && b == nil {
//		return true
//	} else if a != nil && b != nil {
//		if a.Val != b.Val {
//			return false
//		} else {
//			return isOk(a.Left, b.Right) && isOk(a.Right, b.Left)
//		}
//	} else {
//		return false
//	}
//}
