//package main
//
//import "fmt"
//
//func isSubsequence(s string, t string) bool {
//	m, n := len(s), len(t)
//	if m > n {
//		return false
//	}
//	dp := make([][]bool, n+1)
//	for i := 0; i <= n; i++ {
//		dp[i] = make([]bool, m+1)
//		dp[i][0] = true
//	}
//	for i := 1; i <= n; i++ {
//		for j := 1; j <= m; j++ {
//			if s[j-1] == t[i-1] {
//				dp[i][j] = dp[i-1][j] || dp[i-1][j-1]
//			} else {
//				dp[i][j] = dp[i-1][j]
//			}
//		}
//	}
//	return dp[n][m]
//}
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
//func min(a, b int) int {
//	if a > b {
//		return b
//	}
//	return a
//}
//func main() {
//	fmt.Println(isSubsequence("axc", "ahbgxdc"))
//}
