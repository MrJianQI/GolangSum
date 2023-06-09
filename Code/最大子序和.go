//package main
//
//import "fmt"
//
//func maxSubArray(nums []int) int {
//	m := len(nums)
//	res := nums[0]
//	dp := make([]int, m)
//	dp[0] = nums[0]
//	for i := 1; i < m; i++ {
//		if dp[i-1] >= 0 {
//			dp[i] = dp[i-1] + nums[i]
//		} else {
//			dp[i] = nums[i]
//		}
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
//func min(a, b int) int {
//	if a > b {
//		return b
//	}
//	return a
//}
//
//func main() {
//	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
//}
