//package main
//
//import "fmt"
//
//func longestPalindrome(s string) string {
//	m := len(s)
//	dp := make([][]int, m)
//	for i := 0; i < m; i++ {
//		dp[i] = make([]int, m)
//		dp[i][i] = 1
//	}
//	res := string(s[0])
//	temp := 1
//	for i := m - 1; i >= 0; i-- {
//		for j := i; j < m-1; j++ {
//			if i > 0 && dp[i][j] > 0 && s[i-1] == s[j+1] {
//				dp[i-1][j+1] = max(dp[i][j]+2, dp[i-1][j+1])
//				if dp[i-1][j+1] > temp {
//					res = s[i-1 : j+2]
//					temp = dp[i-1][j+1]
//				}
//			}
//			if i+1 < m && (dp[i+1][j] > 0 || i+1 > j) && s[i] == s[j+1] {
//				dp[i][j+1] = max(dp[i][j]+1, dp[i][j+1])
//				if dp[i][j+1] > temp {
//					res = s[i : j+2]
//					temp = dp[i][j+1]
//				}
//			}
//
//		}
//	}
//	//for _, v := range dp {
//	//	fmt.Println(v)
//	//}
//	return res
//}
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
//func main() {
//	fmt.Println(longestPalindrome("babaddabab"))
//}
