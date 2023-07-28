//package main
//
//import (
//	"fmt"
//	"sort"
//)
//
//func threeSum(nums []int) [][]int {
//	sort.Ints(nums)
//	res := make([][]int, 0)
//	n := len(nums)
//	for i := 0; i < n-2; i++ {
//		if i > 1 && nums[i-1] == nums[i] {
//			continue
//		}
//		l, r := i+1, n-1
//		for l < r {
//			left, right := nums[l], nums[r]
//			temp := left + right + nums[i]
//			if temp == 0 {
//				res = append(res, []int{nums[i], left, right})
//				for l < r && left == nums[l] {
//					l++
//				}
//				for l < r && right == nums[r] {
//					r--
//				}
//			} else if temp > 0 {
//				r--
//			} else {
//				l++
//			}
//		}
//	}
//	return res
//}
//func main() {
//	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
//}
