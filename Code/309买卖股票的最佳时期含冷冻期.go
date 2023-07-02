//package main
//
//import "fmt"
//
//func maxProfit(prices []int) int {
//	n := len(prices)
//	if n == 0 {
//		return 0
//	}
//	dp := make([][]int, n)
//	for i := 0; i < n; i++ {
//		dp[i] = make([]int, 3)
//	}
//	dp[0][0] = -prices[0] // 有股票
//	//dp[0][1] = -prices[0] # 无股票，无冷却
//	//dp[0][2] = -prices[0] # 无股票，有冷却
//	for i := 1; i < n; i++ {
//		dp[i][0] = max(dp[i-1][0], max(dp[i-1][2]-prices[i-1], dp[i-1][1]-prices[i]))
//		dp[i][1] = max(dp[i-1][1], dp[i-1][2])
//		dp[i][2] = dp[i-1][0] + prices[i]
//	}
//	return max(dp[n-1][1], dp[n-1][2])
//}
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
//func main() {
//	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
//}
