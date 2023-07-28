//package main
//
//import "fmt"
//
//func trap(height []int) int {
//	n := len(height)
//	if n <= 2 {
//		return 0
//	}
//	left, right := make([]int, n), make([]int, n)
//	res := 0
//	for i, _ := range height {
//		if i == 0 {
//			continue
//		}
//		left[i] = max(left[i-1], height[i-1])
//	}
//	for i := n - 2; i >= 0; i-- {
//		right[i] = max(right[i+1], height[i+1])
//	}
//	fmt.Println(left)
//	fmt.Println(right)
//	for i := 1; i < n-1; i++ {
//		res += max(min(left[i], right[i]), height[i]) - height[i]
//	}
//	return res
//}
//func max(i, j int) int {
//	if i < j {
//		return j
//	}
//	return i
//}
//func min(i, j int) int {
//	if i < j {
//		return i
//	}
//	return j
//}
//func main() {
//	fmt.Println(trap([]int{0, 1, 0, 2}))
//}
