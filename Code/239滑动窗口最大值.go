//package main
//
//import "fmt"
//
//func maxSlidingWindow(nums []int, k int) []int {
//	if k == 1 {
//		return nums
//	}
//	helper := make([]int, 0)
//	res := make([]int, 0)
//	i, j := 0, 0
//	for i < len(nums) {
//		if j-i < k {
//			if j >= len(nums) {
//				break
//			}
//			for len(helper) > 0 && helper[len(helper)-1] < nums[j] {
//				helper = helper[:len(helper)-1]
//			}
//			helper = append(helper, nums[j])
//			j++
//			continue
//		} else {
//			res = append(res, helper[0])
//			if nums[i] == helper[0] {
//				helper = helper[1:]
//			}
//			i++
//		}
//	}
//	return res
//}
//func max(nums []int) int {
//	res := nums[0]
//	for _, v := range nums {
//		if v > res {
//			res = v
//		}
//	}
//	return res
//}
//
//func main() {
//	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3}, 3))
//}
