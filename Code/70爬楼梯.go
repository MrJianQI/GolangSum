//package main
//
//func climbStairs(n int) int {
//	dict := make(map[int]int)
//	dict[0] = 1
//	dict[1] = 1
//	var backtrack func(n int) int
//	backtrack = func(n int) int {
//		if v, ok := dict[n]; ok {
//			return v
//		}
//		temp := backtrack(n-1) + backtrack(n-2)
//		dict[n] = temp
//		return temp
//	}
//	return backtrack(n)
//}
