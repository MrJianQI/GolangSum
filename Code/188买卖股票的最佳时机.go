//package main
//
//import "fmt"
//
//func maxProfit(k int, prices []int) int {
//	m := len(prices)
//	if m == 0 {
//		return 0
//	}
//	dp := make([][]int, m)
//	for i := 0; i < m; i++ {
//		dp[i] = make([]int, 2*k+2)
//		if i == 0 {
//			for j := 0; j < 2*k+2; j += 2 {
//				dp[0][j] = -prices[0]
//			}
//		}
//	}
//
//	//dp[i][0]: 未卖出股票，当前持有股票
//	//dp[i][1]: 未卖出股票，未持有股票
//
//	//dp[i][2]: 卖出1次股票，当前持有股票
//	//dp[i][3]: 卖出1次股票，当前未持有股票
//
//	//dp[i][4]: 卖出2次股票，当前持有股票
//	//dp[i][5]: 卖出2次股票，当前未持有股票
//
//	for i := 1; i < m; i++ {
//		for j := 0; j < k*2+2; j += 2 {
//			if j == 0 {
//				dp[i][j] = max(dp[i-1][j], dp[i-1][j+1]-prices[i])
//				dp[i][j+1] = dp[i-1][j+1]
//			} else {
//				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]-prices[i])
//				dp[i][j+1] = max(dp[i-1][j+1], dp[i-1][j-2]+prices[i])
//			}
//		}
//	}
//	return max(dp[m-1]...)
//}
//
//func max(nums ...int) int {
//	res := nums[0]
//	for _, i := range nums {
//		if i > res {
//			res = i
//		}
//	}
//	return res
//}
//func main() {
//	fmt.Println(maxProfit(2, []int{8, 6, 4, 3, 3, 2, 3, 5, 8, 3, 8, 2, 6}))
//}
