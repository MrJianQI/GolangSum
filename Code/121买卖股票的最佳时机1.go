//package main
//
//func maxProfit(prices []int) int {
//	n := len(prices)
//	if n <= 0 {
//		return 0
//	}
//	res := 0
//	cur := prices[0]
//	for i := 0; i < n; i++ {
//		res = max(res, prices[i]-cur)
//		if prices[i] < cur {
//			cur = prices[i]
//		}
//	}
//	return res
//}
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
