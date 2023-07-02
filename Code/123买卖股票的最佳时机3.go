//package main
//
//func maxProfit(prices []int) int {
//	if len(prices)  == 0{
//		return 0
//	}
//	res := 0
//	n := len(prices)
//	dp := make([][]int,n)
//	for i:=0;i<n;i++{
//		dp[i] = make([]int,5)
//	}
//	dp[0][0] = 0
//	dp[0][1] = -prices[0]
//	dp[0][2] = -9999
//	dp[0][3] = -9999
//	dp[0][4] = -9999
//	for i :=1;i<n;i++{
//		dp[i][0] = dp[i-1][0]
//		dp[i][1] = max(dp[i-1][1],dp[i-1][0] - prices[i])
//		dp[i][2] = max(dp[i-1][2],dp[i-1][1] + prices[i])
//		dp[i][3] = max(dp[i-1][3],dp[i-1][2] - prices[i])
//		dp[i][4] = max(dp[i-1][4],dp[i-1][3] + prices[i])
//	}
//	res = max(res,
//		max(dp[n-1][0],max(dp[n-1][2],dp[n-1][4])),
//	)
//	return res
//}
//func max(a,b int )int{
//	if a > b {
//		return a
//	}
//	return b
//}