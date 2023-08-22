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
//const MAX_N = 1e3 + 5
//const INF = 100000000
//
//var nums [MAX_N][MAX_N]int
//var dp [MAX_N][MAX_N]int
//
//func main() {
//	scanner := bufio.NewScanner(os.Stdin)
//	var n, m int
//	if scanner.Scan() {
//		input := strings.Split(scanner.Text(), " ")
//		n, _ = strconv.Atoi(input[0])
//		m, _ = strconv.Atoi(input[1])
//		for i := 1; i <= n; i++ {
//			if scanner.Scan() {
//				input = strings.Split(scanner.Text(), " ")
//				for j := 1; j <= m; j++ {
//					nums[i][j], _ = strconv.Atoi(input[j-1])
//					dp[i][j] = dp[i-1][j] + dp[i][j-1] + nums[i][j] - dp[i-1][j-1]
//				}
//
//			}
//		}
//	}
//	//for i := 0; i <= n; i++ {
//	//
//	//	fmt.Println(dp[i][:m+1])
//	//}
//	a, b := INF, INF
//	for i := 1; i <= n; i++ {
//		a = min(a, abs(dp[i][m]<<1-dp[n][m]))
//	}
//	for j := 1; j <= m; j++ {
//		b = min(b, abs(dp[n][j]<<1-dp[n][m]))
//	}
//	fmt.Println(min(a, b))
//}
//func abs(n int) int {
//	if n > 0 {
//		return n
//	}
//	return -n
//}
//func min(a, b int) int {
//	if a > b {
//		return b
//	}
//	return a
//}
