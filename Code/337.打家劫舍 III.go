//package main
//func rob(root *TreeNode) int {
//	res := 0
//	var dfs func(root1 *TreeNode) int
//	dict := make(map[*TreeNode]int)
//	dfs = func(root1 *TreeNode)int{
//		if root1 == nil{
//			return 0
//		}
//		if _,ok := dict[root1];ok{
//			return dict[root1]
//		}
//		first := root1.Val
//		sencond := dfs(root1.Left) + dfs(root1.Right)
//		if root1.Left != nil{
//			first += dfs(root1.Left.Right) + dfs(root1.Left.Left)
//		}
//		if root1.Right != nil{
//			first += dfs(root1.Right.Right) + dfs(root1.Right.Left)
//		}
//		dict[root1] = max(first,sencond)
//		// res = max(first,sencond)
//		return dict[root1]
//	}
//	res = dfs(root)
//	return res
//
//}
//func max(a,b int )int{
//	if a > b {
//		return a
//	}
//	return b
//}