//package main
//
//import "fmt"
//
//func minDistance(word1 string, word2 string) int {
//	if len(word1) == 0 && len(word2) == 0 {
//		return 0
//	} else if len(word1) == 0 {
//		return len(word2)
//	} else if len(word2) == 0 {
//		return len(word1)
//	} else {
//		m, n := len(word1), len(word2)
//		dp := make([][]int, m+1)
//		for i := 0; i <= m; i++ {
//			dp[i] = make([]int, n+1)
//			dp[i][0] = i
//		}
//		for i := 0; i <= n; i++ {
//			dp[0][i] = i
//		}
//		for i := 1; i <= m; i++ {
//			for j := 1; j <= n; j++ {
//				if word1[i-1] == word2[j-1] {
//					dp[i][j] = dp[i-1][j-1]
//				} else {
//					dp[i][j] = min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j])) + 1
//				}
//
//			}
//		}
//		//for _, i := range dp {
//		//	fmt.Println(i)
//		//}
//		return dp[m][n]
//	}
//}
//func min(a, b int) int {
//	if a > b {
//		return b
//	}
//	return a
//}
//func main() {
//	fmt.Println(minDistance("h", "r"))
//}
