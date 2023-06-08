//package main
//
//import "fmt"
//
//func numberOfArithmeticSlices(nums []int) int {
//	if len(nums) < 3 {
//		return 0
//	}
//	res := 0
//	dp := make([]int, len(nums))
//	for i := 2; i < len(nums); i++ {
//		if nums[i]+nums[i-2] == nums[i-1]*2 {
//			dp[i] = dp[i-1] + 1
//		} else {
//			dp[i] = 0
//		}
//		res += dp[i]
//	}
//	//fmt.Println(dp)
//	return res
//}
//
//func main() {
//	fmt.Println(numberOfArithmeticSlices([]int{1, 2, 3, 7, 8, 9}))
//}
