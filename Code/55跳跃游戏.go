//package main
//
//import (
//	"fmt"
//)
//
//func canJump(nums []int) bool {
//	temp := 0
//	for i := 0; i < len(nums); i++ {
//		if temp >= len(nums)-1 {
//			return true
//		}
//		if nums[i] == 0 && temp <= i {
//			return false
//		}
//		temp = max(temp, nums[i]+i)
//	}
//	return true
//}
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
//func main() {
//	fmt.Println(canJump([]int{3, 0, 8, 2, 0, 0, 1}))
//}
