//package main
//
//import "fmt"
//
//func generate(numRows int) [][]int {
//	res := make([][]int, 0)
//	res = append(res, []int{1})
//	for i := 1; i < numRows; i++ {
//		temp := make([]int, 0)
//		temp = append(temp, 1)
//
//		for j := 0; j < i-1; j++ {
//			temp = append(temp, res[i-1][j]+res[i-1][j+1])
//		}
//		temp = append(temp, 1)
//		res = append(res, temp)
//	}
//
//	return res
//}
//func main() {
//	fmt.Println(generate(5))
//}
