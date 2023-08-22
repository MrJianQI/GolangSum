//package main
//
//import (
//	"fmt"
//)
//
//func wordBreak(s string, wordDict []string) bool {
//	if len(s) == 0 {
//		return true
//	}
//	dp := make([]bool, len(s)+1)
//	dp[0] = true
//	dict := make(map[string]int)
//	for _, v := range wordDict {
//		dict[v] = 1
//	}
//	var backtrack func(str string)
//	backtrack = func(str string) {
//		for i := 0; i < len(s); i++ {
//			for j := i + 1; j < len(s)+1; j++ {
//				if dp[i] && isok(s[i:j], dict) {
//					dp[j] = true
//				}
//			}
//		}
//	}
//	backtrack(s)
//	return dp[len(s)]
//}
//func isok(s string, dict map[string]int) bool {
//	_, ok := dict[s]
//	return ok
//}
//func main() {
//	fmt.Println(wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}))
//}
//
////func max(a, b int) int {
////	if a > b {
////		return a
////	}
////	return b
////}
////func isOK(s string, dict map[string]int) bool {
////	if s == "" {
////		return true
////	}
////	_, ok := dict[s]
////	return ok
////}
