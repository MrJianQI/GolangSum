//package main
//
//func numSquares(n int) int {
//	var dfs func(s int) int
//	dict := make(map[int]int)
//	dfs = func(s int) int {
//		if s == 0 {
//			return 0
//		}
//		if s == 1 {
//			return 1
//		}
//		if _, ok := dict[s]; ok {
//			return dict[s]
//		}
//		temp := s
//		for i := 1; i*i <= s; i++ {
//			temp = min(temp, dfs(s-i*i)+1)
//		}
//		dict[s] = temp
//		return temp
//	}
//	return dfs(n)
//}
//func min(i, j int) int {
//	if i > j {
//		return j
//	}
//	return i
//}
