//package main
//
//import (
//	"fmt"
//)
//
//func maxUncrossedLines(nums1 []int, nums2 []int) int {
//	m, n := len(nums1), len(nums2)
//	dp := make([][]int, m+1)
//	for i := 0; i <= m; i++ {
//		dp[i] = make([]int, n+1)
//	}
//	for i := 1; i <= m; i++ {
//		for j := 1; j <= n; j++ {
//			if nums1[i-1] == nums2[j-1] {
//				dp[i][j] = dp[i-1][j-1] + 1
//			} else {
//				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
//			}
//		}
//	}
//	return dp[m][n]
//}
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
//func main() {
//	fmt.Println(maxUncrossedLines([]int{1, 3, 7, 1, 7, 5}, []int{1, 9, 2, 5, 1}))
//}
