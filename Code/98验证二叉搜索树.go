//package main
//
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//func isValidBST(root *TreeNode) bool {
//	var dfs func(cur *TreeNode, start, end int) bool
//	dfs = func(cur *TreeNode, start, end int) bool {
//		if start > end {
//			return false
//		}
//		if cur == nil {
//			return true
//		}
//		if cur.Val <= start || cur.Val >= end {
//			return false
//		}
//		left := dfs(cur.Left, start, cur.Val)
//		right := dfs(cur.Right, cur.Val, end)
//		return left && right
//	}
//	return dfs(root, -100000000000, 100000000000000)
//}