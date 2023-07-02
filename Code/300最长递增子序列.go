//package main
//
//import "fmt"
//
//func lengthOfLIS(nums []int) int {
//	m := len(nums)
//	if m == 0 {
//		return 0
//	}
//	dp := make([]int, m+1)
//	res := 1
//	dp[1] = 1
//	for i := 1; i <= m; i++ {
//		for j := 1; j < i; j++ {
//			if nums[i-1] > nums[j-1] {
//				dp[i] = max(dp[i], dp[j]+1)
//			}
//		}
//		dp[i] = max(dp[i], 1)
//		if dp[i] > res {
//			res = dp[i]
//		}
//	}
//
//	return res
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
//func main() {
//	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
//}
