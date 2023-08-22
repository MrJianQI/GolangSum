//package main
//
//import "fmt"
//
//func numIslands(grid [][]byte) int {
//	res := 0
//	m, n := len(grid), len(grid[0])
//	dir := make([][]int, 0)
//	dir = append(dir, []int{0, 1})
//	dir = append(dir, []int{0, -1})
//	dir = append(dir, []int{1, 0})
//	dir = append(dir, []int{-1, 0})
//	visited := make([][]int, m)
//	for i := 0; i < m; i++ {
//		visited[i] = make([]int, n)
//	}
//	var dfs func(i, j int)
//	dfs = func(i, j int) {
//		if i < 0 || i >= m || j < 0 || j >= n {
//			return
//		}
//		if visited[i][j] == 1 || grid[i][j] == '0' {
//			return
//		}
//		visited[i][j] = 1
//		for _, v := range dir {
//			dfs(i+v[0], j+v[1])
//		}
//	}
//	for i := 0; i < m; i++ {
//		for j := 0; j < n; j++ {
//			if grid[i][j] == '1' && visited[i][j] == 0 {
//				res++
//				dfs(i, j)
//
//			}
//		}
//	}
//	return res
//}
//func main() {
//	fmt.Println(numIslands([][]byte{
//		[]byte{'1', '1', '0', '0', '0'},
//		[]byte{'1', '1', '0', '0', '0'},
//		[]byte{'0', '0', '1', '0', '0'},
//		[]byte{'0', '0', '0', '1', '1'},
//	}))
//}
