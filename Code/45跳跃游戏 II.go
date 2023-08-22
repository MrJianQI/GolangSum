//package main
//
//import "fmt"
//
//func jump(nums []int) int {
//	n := len(nums)
//	res := 0
//	start, end := 0, 1
//	for end < n {
//		a := 0
//		for i := start; i < end; i++ {
//			a = max(a, i+nums[i])
//		}
//		start = end
//		end = a + 1
//		res++
//	}
//
//	return res
//}
//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
//func main() {
//	fmt.Println(jump([]int{2, 1}))
//}
