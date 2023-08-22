package main

import "fmt"

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	visited := make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}
	dir := make([][]int, 4)
	dir[0] = []int{1, 0}
	dir[1] = []int{-1, 0}
	dir[2] = []int{0, 1}
	dir[3] = []int{0, -1}
	var dfs func(i, j int, str string) bool
	dfs = func(i, j int, str string) bool {
		if len(str) == 0 {
			return true
		}
		if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] == 1 {
			return false
		}
		if board[i][j] != str[0] {
			return false
		}
		visited[i][j] = 1
		for k := 0; k < 4; k++ {
			if dfs(i+dir[k][0], j+dir[k][1], str[1:]) {
				return true
			}
		}
		visited[i][j] = 0
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				if dfs(i, j, word) {
					return true
				}
			}
		}
	}
	return false
}
func main() {
	board := [][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println(exist(board, word))
}
