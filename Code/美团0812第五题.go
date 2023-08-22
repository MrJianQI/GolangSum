//package main
//
//import "fmt"
//
//type LinkNode struct {
//	Val   int
//	Left  *LinkNode
//	Right *LinkNode
//}
//
//func main() {
//	n := 0
//	fmt.Scan(&n)
//	nodes := make([]*LinkNode, n)
//	for i := 0; i < n; i++ {
//		fmt.Scan(&nodes[i].Val)
//	}
//	for i := 0; i < n-1; i++ {
//		x, y := 0, 0
//		fmt.Scan(&x, &y)
//		if x < y {
//			x, y = y, x
//		}
//		nodes[x-1].Right = nodes[y-1]
//		nodes[y-1].Left = nodes[x-1]
//	}
//	res := 0
//	dp := make([][]int, n)
//	for i := 0; i < n; i++ {
//		dp[i] = make([]int, 2)
//	}
//	var dfs func(cur *LinkNode) int
//	dfs = func(cur *LinkNode) int {
//		if cur == nil {
//			return 0
//		}
//		if cur.Left != nil && isOK(cur.Left.Val*cur.Val) {
//			a := 2 + dfs(cur.Left.Left) + dfs(cur.Left.Right)
//			b := dfs(cur.Left)
//			return max(a, b)
//		}
//		if cur.Right != nil && isOK(cur.Right.Val*cur.Val) {
//			a := 2 + dfs(cur.Right.Left) + dfs(cur.Right.Right)
//			b := dfs(cur.Right)
//			return max(a, b)
//		}
//		return 0
//	}
//}
//func isOK(num int) bool {
//	if num == 1 {
//		return true
//	}
//	i := 0
//	for i*i <= num {
//		if i*i == num {
//			return true
//		}
//		i++
//	}
//	return false
//}
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
