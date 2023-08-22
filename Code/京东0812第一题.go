//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strconv"
//	"strings"
//)
//
//const MAX_N = 1005
//const INF = 10000000
//
//var dp [MAX_N][MAX_N]int
//
//func main() {
//	scanner := bufio.NewScanner(os.Stdin)
//	n := 0
//	s := ""
//	if scanner.Scan() {
//		input := strings.Split(scanner.Text(), " ")
//		n, _ = strconv.Atoi(input[0])
//		if scanner.Scan() {
//			s = scanner.Text()
//		}
//	}
//	for i := 0; i < n; i++ {
//		for j := 0; j < n; j++ {
//			if i == j {
//				dp[i][j] = 0
//				continue
//			}
//			dp[i][j] = INF
//		}
//	}
//	for i := n - 2; i >= 0; i-- {
//		for j := i + 1; j < n; j++ {
//			if s[i] == s[j] && ((i+1 == j-1) || (i+1 == j) || (i+1 < n && j-1 >= 0 && dp[i+1][j-1] > 0)) {
//				dp[i][j] = min(dp[i][j], dp[i+1][j-1])
//			}
//			if s[i] != s[j] && ((i+1 == j-1) || (i+1 == j) || (i+1 < n && j-1 >= 0 && dp[i+1][j-1] > 0)) {
//				dp[i][j] = dp[i+1][j-1] + 1
//			}
//			if i+1 < j && s[i] == s[i+1] && isOk(s) {
//				dp[i][j] = dp[i+1][j] + 1
//			}
//		}
//	}
//	fmt.Println(dp[0][n-1])
//}
//func isOk(s string) bool {
//	if len(s) == 1 {
//		return true
//	}
//	i, j := 0, len(s)-1
//	for i < j {
//		if s[i] != s[j] {
//			return false
//		}
//	}
//	return true
//}
//func min(a, b int) int {
//	if a > b {
//		return b
//	}
//	return b
//}
