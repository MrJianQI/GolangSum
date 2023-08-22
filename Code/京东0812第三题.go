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
//const MAX_N = 50 + 1
//const INF = 10000000
//
//var dp [MAX_N][MAX_N]int
//
//func main() {
//	scanner := bufio.NewScanner(os.Stdin)
//	var n, m int
//	if scanner.Scan() {
//		inputNums := strings.Split(scanner.Text(), " ")
//		n, _ = strconv.Atoi(inputNums[0])
//		m, _ = strconv.Atoi(inputNums[1])
//		for i := 0; i < n; i++ {
//			if scanner.Scan() {
//				for index, v := range scanner.Text() {
//					if v == '.' {
//						dp[i][index] = 0
//					} else {
//						dp[i][index] = 1
//					}
//				}
//			}
//		}
//	}
//	res := 0
//	for i := 0; i < n; i++ {
//		for j := 0; j < m; j++ {
//			if dp[i][j] == 0 {
//				continue
//			}
//			// 向下
//			for k := i + 1; k < n; k++ {
//				if dp[k][j] == 1 {
//					length := k - i
//					if j+length < n && dp[i][j+length] == 1 && dp[k][j+length] == 1 {
//						//fmt.Println(i, j, k, j, i, j+length, k, j+length)
//						res++
//					}
//				}
//			}
//			// 斜方向上
//			for x := i + 1; x < n; x++ {
//				for y := 0; y < j; y++ {
//					if dp[x][y] != 1 {
//						continue
//					}
//					dx, dy := abs(i-x), abs(j-y)
//					if x+dy < n && y+dx < m && dp[x+dy][y+dx] == 1 &&
//						i+dy < n && j+dx < m && dp[i+dy][j+dx] == 1 {
//						//fmt.Println("A", i, j, x, y, x+dy, y+dx, i+dy, j+dx)
//						res++
//					}
//				}
//
//			}
//
//		}
//	}
//	fmt.Println(res)
//}
//func abs(n int) int {
//	if n < 0 {
//		return -n
//	}
//	return n
//}
