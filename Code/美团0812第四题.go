//package main
//
//import "fmt"
//
//func main() {
//	var length int
//	var s string
//	fmt.Scan(&length)
//	fmt.Scan(&s)
//	list := make([][]int, 0)
//	for i := 1; i < length; i++ {
//		if length%i == 0 {
//			list = append(list, []int{i, length / i})
//		}
//	}
//	res := length
//	for _, v := range list {
//		temp := 0
//		m, n := v[0], v[1]
//		nums := make([][]uint8, m)
//		visited := make([][]int, m)
//		for i := 0; i < m; i++ {
//			nums[i] = make([]uint8, n)
//			visited[i] = make([]int, n)
//		}
//		for i := 0; i < m; i++ {
//			for j := 0; j < n; j++ {
//				nums[i][j] = s[i*n+j]
//			}
//		}
//		var dfs func(i, j int)
//		dfs = func(x, y int) {
//			if x < 0 || x >= m || y < 0 || y >= n || visited[x][y] == 1 {
//				return
//			}
//			visited[x][y] = 1
//			for _, v := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
//				dx, dy := v[0], v[1]
//				nx, ny := x+dx, y+dy
//				if nx < 0 || nx >= m || ny < 0 || ny >= n {
//					continue
//				}
//				if nums[nx][ny] == nums[x][y] {
//					dfs(nx, ny)
//				}
//			}
//		}
//		for i := 0; i < m; i++ {
//			for j := 0; j < n; j++ {
//				if visited[i][j] == 0 {
//					temp++
//					dfs(i, j)
//				}
//			}
//		}
//		res = min(res, temp)
//	}
//	fmt.Println(res)
//}
//func min(a, b int) int {
//	if a > b {
//		return b
//	}
//	return a
//}
