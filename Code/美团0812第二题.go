//package main
//
//import "fmt"
//
//func main() {
//	var n int
//	fmt.Scan(&n)
//	nums := make([]int, n)
//	left := make([]int, n)
//	right := make([]int, n)
//
//	for i := 0; i < n; i++ {
//		fmt.Scan(&nums[i])
//		if i == 0 {
//			left[i] = nums[i]
//		} else {
//			left[i] = left[i-1] + nums[i]
//		}
//	}
//	for i := n - 1; i >= 0; i-- {
//		if i == n-1 {
//			right[n-i-1] = nums[i]
//		} else {
//			right[n-i-1] = right[n-i-2] + nums[i]
//		}
//	}
//	var x, y int
//	fmt.Scan(&x, &y)
//	l := left[y-1]
//	if y > 1 {
//		l = left[y-1] - left[y-2]
//	}
//	r := right[x-1]
//	if x > 1 {
//		r = right[x-1] - right[x-2]
//	}
//	fmt.Println(min(l, r))
//}
//func min(a, b int) int {
//	if a > b {
//		return b
//	}
//	return a
//}
