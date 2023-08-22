//package main
//
//import "fmt"
//
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//func rightSideView(root *TreeNode) []int {
//	res := make([]int, 0)
//	stack1111 := make([]*TreeNode, 0)
//	stack1111 = append(stack1111, root)
//	for len(stack1111) > 0 {
//		n := len(stack1111)
//		a := 0
//		for i := 0; i < n; i++ {
//			a = stack1111[i].Val
//			if stack1111[i].Left != nil {
//				stack1111 = append(stack1111, stack1111[i].Left)
//			}
//			if stack1111[i].Right != nil {
//				stack1111 = append(stack1111, stack1111[i].Right)
//			}
//		}
//		res = append(res, a)
//		stack1111 = stack1111[n:]
//	}
//	return res
//}
//func main() {
//	root1 := &TreeNode{Val: 1}
//	root2 := &TreeNode{Val: 2}
//	root3 := &TreeNode{Val: 3}
//	root4 := &TreeNode{Val: 4}
//	root5 := &TreeNode{Val: 5}
//	root1.Left = root2
//	root1.Right = root3
//	root2.Right = root5
//	root3.Right = root4
//	fmt.Println(rightSideView(root1))
//}
