//package main
//
//import (
//	"fmt"
//	"sort"
//)
//
//func coinChange(coins []int, amount int) int {
//	if amount == 0 {
//		return 0
//	}
//	sort.Slice(coins, func(i, j int) bool {
//		return coins[i] < coins[j]
//	})
//	n := len(coins)
//	temp := make([]int, 0)
//	for i := 0; i <= n; i++ {
//		temp = append(temp, amount+1)
//	}
//	dp := make([][]int, amount+1)
//	for i := 0; i <= amount; i++ {
//		dp[i] = append(dp[i], temp...)
//	}
//	for i := 0; i <= n; i++ {
//		dp[0][i] = 0
//	}
//	for i := 1; i < amount+1; i++ {
//		for j := 1; j <= n; j++ {
//			if i < coins[j-1] {
//				dp[i][j] = dp[i][j-1]
//			} else {
//				dp[i][j] = min(dp[i][j-1], dp[i-coins[j-1]][j]+1)
//			}
//
//		}
//	}
//	if dp[amount][n] == amount+1 {
//		return 0
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
