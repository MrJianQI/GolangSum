//package main
//func rob(nums []int) int {
//	n := len(nums)
//	if n <= 0 {
//		return 0
//	}
//	var dpp func (datas []int)int
//	dpp = func(datas []int)int{
//		if len(datas) <= 0{
//			return 0
//		}
//		if len(datas) == 1 {
//			return datas[0]
//		}
//		m := len(datas)
//		dp := make([]int,m )
//		dp[0] = datas[0]
//		dp[1] = max(datas[0],datas[1])
//		for i:=2;i<m;i++{
//			dp[i] = max(dp[i-2] + datas[i],dp[i-1])
//		}
//		return dp[m-1]
//	}
//	return max(nums[0],max(dpp(nums[:n-1]),dpp(nums[1:])))
//}
//func max(a,b int )int{
//	if a > b {
//		return a
//	}
//	return b
//}