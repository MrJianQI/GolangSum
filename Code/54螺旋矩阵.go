//package main
//
//import "fmt"
//
//
//func spiralOrder(matrix [][]int) []int {
//	n := len(matrix)
//	m := len(matrix[0])
//	up, down, left, right := 0, n-1, 0, m-1
//	res := make([]int, 0)
//	for up <= down && left <= right {
//		for i := left; i <= right; i++ {
//			res = append(res, matrix[up][i])
//		}
//		up++
//		for i := up; i <= down; i++ {
//			res = append(res, matrix[i][right])
//		}
//		right--
//		if up > down || left > right {
//			break
//		}
//		for i := right; i >= left; i-- {
//			res = append(res, matrix[down][i])
//		}
//		down--
//		for i := down; i >= up; i-- {
//			res = append(res, matrix[i][left])
//		}
//		left++
//	}
//	return res
//}
//
//func main() {
//	fmt.Println(spiralOrder([][]int{
//		[]int{1, 2, 3, 4},
//		[]int{5, 6, 7, 8},
//		[]int{9, 10, 11, 12},
//	}))
//}
