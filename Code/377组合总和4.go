//package main
//
//import (
//	"fmt"
//	"sort"
//)
//
//func combinationSum4(nums []int, target int) int {
//	n := len(nums)
//	sort.Ints(nums)
//	dp := make([][]int, target+1)
//	for i := 0; i <= target; i++ {
//		dp[i] = make([]int, n+1)
//	}
//	for i := 0; i <= n; i++ {
//		dp[0][i] = 1
//	}
//	for i := 1; i <= target; i++ {
//		for j := 1; j < n; j++ {
//			if i < nums[j] {
//				dp[i][j] = dp[i][j] + dp[i][j-1]
//			} else {
//				dp[i][j] = dp[i][j] + dp[i-nums[j-1]][j] + dp[i][j-1]
//			}
//		}
//	}
//	for _, v := range dp {
//		fmt.Println(v)
//	}
//	return dp[target][n]
//}
//func main() {
//	fmt.Println(combinationSum4([]int{1, 2, 3}, 32))
//}
