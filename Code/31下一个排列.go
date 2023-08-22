//package main
//
//import (
//	"sort"
//)
//
//func nextPermutation(nums []int) {
//	if len(nums) == 1 {
//		return
//	}
//	if maxtomin(nums) {
//		sort.Ints(nums)
//		return
//	}
//	if maxtomin(nums[1:]) {
//		v, ok := ismax(nums[1:], nums[0])
//		if ok == false {
//			sort.Ints(nums)
//			return
//		} else {
//			nums[v+1], nums[0] = nums[0], nums[v+1]
//			sort.Ints(nums[1:])
//			return
//		}
//	} else {
//		nextPermutation(nums[1:])
//		//index := isdec(nums[1:])
//		//nums[index+1], nums[index] = nums[index], nums[index+1]
//		//sort.Ints(nums[index+1:])
//		return
//	}
//}
//func maxtomin(nums []int) bool {
//	for i := 0; i < len(nums)-1; i++ {
//		if nums[i] < nums[i+1] {
//			return false
//		}
//	}
//	return true
//}
//func ismax(nums []int, target int) (int, bool) {
//	for i := len(nums) - 1; i >= 0; i-- {
//		if nums[i] > target {
//			return i, true
//		}
//	}
//	return -1, false
//}
//func isdec(nums []int) int {
//	for i := len(nums) - 1; i >= 0; i-- {
//		if nums[i] > nums[i-1] {
//			return i
//		}
//	}
//	return -1
//}
//func main() {
//	nextPermutation([]int{5, 4, 7, 5, 3, 2})
//}
