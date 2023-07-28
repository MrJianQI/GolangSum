//package main
//
//import (
//	"fmt"
//	"sort"
//)
//
//func coinChange(coins []int, amount int) int {
//	n := len(coins)
//	dp := make([][]int, amount+1)
//	sort.Slice(coins, func(i, j int) bool {
//		return coins[i] < coins[j]
//	})
//	for i := 0; i <= amount; i++ {
//		dp[i] = make([]int, n+1)
//		for j := 0; j <= n; j++ {
//			dp[i][j] = amount + 2
//		}
//
//	}
//	for i := 0; i <= n; i++ {
//		dp[0][i] = 0
//	}
//
//	for i := 1; i <= amount; i++ {
//		for j := 1; j <= n; j++ {
//			if i < coins[j-1] {
//				dp[i][j] = dp[i][j-1]
//			} else {
//				dp[i][j] = min(dp[i][j-1], dp[i-coins[j-1]][j]+1)
//
//			}
//		}
//
//	}
//	for _, v := range dp {
//		fmt.Println(v)
//	}
//	if dp[amount][n] == amount+2 {
//		return -1
//	}
//	return dp[amount][n]
//}
//
//func min(a, b int) int {
//	if a > b {
//		return b
//	}
//	return a
//}
//
//func main() {
//	fmt.Println(coinChange([]int{1, 2, 5}, 5))
//}
