//package main
//
//import "fmt"
//
//func subarraySum(nums []int, k int) int {
//	res := 0
//	dp := make([]int, 0)
//	dp[0] = nums[0]
//	for i := 1; i < len(nums); i++ {
//		dp[i] = dp[i-1] + nums[i]
//	}
//	dict := make(map[int]int)
//	dict[0] = 1
//	for i := 0; i < len(nums); i++ {
//		if v, ok := dict[dp[i]-k]; ok {
//			res = res + v
//		}
//		dict[dp[i]]++
//	}
//	return res
//}
//func main() {
//	fmt.Println(subarraySum([]int{-1, -1, 1, -1, -1}, -1))
//}
