//package main
//
//func productExceptSelf(nums []int) []int {
//	n := len(nums)
//	left, right := make([]int, n), make([]int, n)
//	res := make([]int, n)
//	left[0] = 1
//	right[n-1] = 1
//	for i := n - 2; i >= 0; i-- {
//		right[i] = right[i+1] * nums[i+1]
//	}
//	for i := 0; i < n; i++ {
//		if i > 0 {
//			left[i] = left[i-1] * nums[i-1]
//		}
//		res[i] = left[i] * right[i]
//	}
//	return res
//}
