//package main
//
//import (
//	"fmt"
//	"sort"
//)
//
//func canPartition(nums []int) bool {
//	n := len(nums)
//	if n == 0 {
//		return false
//	}
//	m := sum(nums)
//	if m%2 == 1 {
//		return false
//	}
//	sort.Ints(nums)
//	dp := make([][]int, m/2+1)
//	temp := make([]int, n+1)
//	for i := 0; i <= n; i++ {
//		temp[i] = m/2 + 1
//	}
//	for i := 0; i <= m/2; i++ {
//		dp[i] = append(dp[i], temp...)
//	}
//	for i := 0; i <= n; i++ {
//		dp[0][i] = 0
//
//	}
//	for i := 1; i <= m/2; i++ {
//		for j := 1; j <= n; j++ {
//			if i >= nums[j-1] {
//				dp[i][j] = min(dp[i][j-1], dp[i-nums[j-1]][j-1]+1)
//			} else {
//				dp[i][j] = dp[i][j-1]
//			}
//		}
//	}
//
//	if dp[m/2][n] == m/2+1 {
//		return false
//	}
//	return true
//}
//func sum(nums []int) int {
//	res := 0
//	for _, v := range nums {
//		res += v
//	}
//	return res
//}
//func min(a, b int) int {
//	if a > b {
//		return b
//	}
//	return a
//}
//func main() {
//	fmt.Println(canPartition([]int{4, 4, 3, 1}))
//}
