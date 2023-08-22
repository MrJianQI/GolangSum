//package main
//
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//func flatten(root *TreeNode) {
//	if root == nil {
//		return
//	}
//	flatten(root.Left)
//	flatten(root.Right)
//	left, right := root.Left, root.Right
//	root.Left = nil
//	root.Right = left
//	cur := root
//	for cur.Right != nil {
//		cur = cur.Right
//	}
//	cur.Right = right
//	return
//}
