//package main
//
//import "fmt"
//
//func longestPalindromeSubseq(s string) int {
//	if len(s) <= 1 {
//		return len(s)
//	}
//	res := 0
//	dp := make([][]int, len(s))
//	for i := 0; i < len(s); i++ {
//		dp[i] = make([]int, len(s))
//		dp[i][i] = 1
//
//	}
//	for i := len(s) - 1; i >= 0; i-- {
//		for j := i + 1; j < len(s); j++ {
//			if s[i] == s[j] {
//				dp[i][j] = dp[i+1][j-1] + 2
//			} else {
//				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
//			}
//			if dp[i][j] > res {
//				res = dp[i][j]
//			}
//		}
//	}
//	return res
//}
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
//
//func main() {
//	fmt.Println(longestPalindromeSubseq("bbbab"))
//}
