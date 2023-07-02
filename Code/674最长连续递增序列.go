//package main
//
//import "fmt"
//
//func findLengthOfLCIS(nums []int) int {
//	m := len(nums)
//	if m <= 0 {
//		return 0
//	}
//	dp := make([]int, m+1)
//	res := 1
//	dp[1] = 1
//	for i := 2; i <= m; i++ {
//		if nums[i-1] > nums[i-2] {
//			dp[i] = dp[i-1] + 1
//		} else {
//			dp[i] = 1
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
//	if a < b {
//		return b
//	}
//	return a
//}
//func main() {
//	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 7}))
//}
