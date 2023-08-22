//package main
//
//func diameterOfBinaryTree(root *TreeNode) int {
//	// if root == nil {
//	// 	return 0
//	// }
//	res := 0
//	var maxLen func(cur *TreeNode) int
//	maxLen = func(cur *TreeNode) int {
//		if cur == nil {
//			return 0
//		}
//		if cur.Left == nil && cur.Right == nil {
//			return 1
//		}
//		left := maxLen(cur.Left)
//		right := maxLen(cur.Right)
//		res = max(res, left+right)
//
//		return max(left, right) + 1
//	}
//	maxLen(root)
//	return res
//}
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
