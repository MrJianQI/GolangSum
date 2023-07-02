//package main
//
//func findLength(nums1 []int, nums2 []int) int {
//	m, n := len(nums1), len(nums2)
//	if m == 0 || n == 0 {
//		return 0
//	}
//	dp := make([][]int, m+1)
//	for i := 0; i <= m; i++ {
//		dp[i] = make([]int, n+1)
//	}
//	res := 0
//	for i := 1; i <= m; i++ {
//		for j := 1; j <= n; j++ {
//			if nums1[i-1] == nums2[j-1] {
//				if i >= 2 && j >= 2 && nums1[i-2] == nums2[j-2] {
//					dp[i][j] = dp[i-1][j-1] + 1
//				} else {
//					dp[i][j] = 1
//				}
//			} else {
//				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
//			}
//			if dp[i][j] > res {
//				res = dp[i][j]
//			}
//		}
//	}
//	return res
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
