//package main
//
//func maxProduct(nums []int) int {
//
//	n := len(nums)
//	dp1 := make([]int, n)
//	dp2 := make([]int, n)
//
//	dp1[0] = nums[0]
//	dp2[0] = nums[0]
//	res := nums[0]
//	for i := 1; i < n; i++ {
//		dp1[i] = max(max(dp1[i-1]*nums[i], dp2[i-1]*nums[i]), nums[i])
//		dp2[i] = min(nums[i], min(dp1[i-1]*nums[i], dp2[i-1]*nums[i]))
//		if dp1[i] > res {
//			res = dp1[i]
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
//func min(a, b int) int {
//	if a > b {
//		return b
//	}
//	return a
//}
