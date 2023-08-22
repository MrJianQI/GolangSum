//package main
//
//func generateParenthesis(n int) []string {
//	res := make([]string, 0)
//
//	var dfs func(left, right int, cur string)
//	dfs = func(left, right int, cur string) {
//		if left > right {
//			return
//		}
//		if left == right {
//			if left == 0 {
//				res = append(res, cur)
//				return
//			} else {
//				left--
//				dfs(left, right, cur+"(")
//				left++
//			}
//		} else {
//			left--
//			dfs(left, right, cur+"(")
//			left++
//			right--
//			dfs(left, right, cur+")")
//			right++
//		}
//	}
//	dfs(n, n, "")
//	return res
//}
