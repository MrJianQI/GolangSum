//package main
//func maxProfit(prices []int) int {
//	if len(prices) <= 1 {
//		return 0
//	}
//	res := 0
//	n := len(prices)
//	dp := make([][]int,n)
//	for i:=0;i<n;i++{
//		dp[i] = make([]int,2)
//	}
//	dp[0][0] = 0
//	dp[0][1] = -prices[0]
//	for i:=1;i<n;i++{
//		dp[i][0] = max(dp[i-1][0],dp[i-1][1]+ prices[i])
//		dp[i][1] = max(dp[i-1][1],dp[i-1][0]- prices[i])
//	}
//	res = max(res,dp[n-1][0])
//	return res
//}
//func max(a,b int )int{
//	if a > b {
//		return a
//	}
//	return b
//}