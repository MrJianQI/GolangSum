//package main
//
//import "sort"
//
//func longestConsecutive(nums []int) int {
//	sort.Ints(nums)
//	dp := make([]int, len(nums))
//	dp[0] = 1
//	max := 1
//	for i := 1; i < len(nums); i++ {
//		if nums[i] == nums[i-1]+1 {
//			dp[i] = dp[i-1] + 1
//			if max < dp[i] {
//				max = dp[i]
//			}
//		} else {
//			dp[i] = 1
//		}
//	}
//	return max
//}
