//package main
//
//import "fmt"
//
//func orangesRotting(grid [][]int) int {
//	m, n := len(grid), len(grid[0])
//	dir := make([][]int, 0)
//	dir = append(dir, []int{0, 1})
//	dir = append(dir, []int{0, -1})
//	dir = append(dir, []int{1, 0})
//	dir = append(dir, []int{-1, 0})
//	visited := make([][]int, 0)
//	for i := 0; i < m; i++ {
//		visited = append(visited, make([]int, n))
//
//	}
//	res := 0
//	var dfs func(i, j int) bool
//	dfs = func(i, j int) bool {
//		if i < 0 || i >= m || j < 0 || j >= n {
//			return false
//		}
//		if visited[i][j] == 1 || grid[i][j] == 0 || grid[i][j] == 2 {
//			return false
//		}
//		grid[i][j] = 2
//		return true
//	}
//	list := make([][]int, 0)
//	for i := 0; i < m; i++ {
//		for j := 0; j < n; j++ {
//			if grid[i][j] == 2 {
//				list = append(list, []int{i, j})
//			}
//		}
//	}
//	for len(list) != 0 {
//		//res++
//		temp := make([][]int, 0)
//		for _, v := range grid {
//			fmt.Println(v)
//		}
//		fmt.Println("______________________")
//		for _, v1 := range list {
//			for _, v2 := range dir {
//				if dfs(v1[0]+v2[0], v1[1]+v2[1]) {
//					temp = append(temp, []int{v1[0] + v2[0], v1[1] + v2[1]})
//				}
//			}
//		}
//		//if len(temp) != 0 {
//		res++
//		//}
//		list = temp
//	}
//	if res == 0 {
//		return -1
//	}
//	return res - 1
//}
//func min(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
//func main() {
//	fmt.Println(orangesRotting([][]int{
//		[]int{2, 1, 1},
//		[]int{1, 1, 0},
//		[]int{0, 1, 1},
//	}))
//}
