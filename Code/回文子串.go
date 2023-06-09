//package main
//
//import "fmt"
//
//func countSubstrings(s string) int {
//	if len(s) <= 1 {
//		return len(s)
//	}
//	dp := make([][]bool, len(s))
//	for i := 0; i < len(s); i++ {
//		dp[i] = make([]bool, len(s))
//		dp[i][i] = true
//	}
//	res := len(s)
//	for i := len(s) - 1; i >= 0; i-- {
//		for j := i + 1; j < len(s); j++ {
//			if s[i] == s[j] && (i == j-1 || dp[i+1][j-1] == true) {
//				dp[i][j] = true
//				res++
//			}
//		}
//	}
//	return res
//}
//
//func main() {
//	fmt.Println(countSubstrings("abcabc"))
//}
