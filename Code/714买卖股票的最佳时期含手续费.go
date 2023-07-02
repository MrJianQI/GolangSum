//package main
//
//import "fmt"
//
//func maxProfit(prices []int, fee int) int {
//	n := len(prices)
//	if n == 0 {
//		return 0
//	}
//	dp := make([][]int, n)
//	for i := 0; i < n; i++ {
//		dp[i] = make([]int, 2)
//	}
//	dp[0][0] = -prices[0]
//	for i := 1; i < n; i++ {
//		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
//		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]-fee)
//	}
//	return dp[n-1][1]
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
//func main() {
//	fmt.Println(maxProfit([]int{1, 3, 7, 5, 10, 3}, 3))
//}
