//package main
//
//func moveZeroes(nums []int) {
//	if len(nums) <= 1 {
//		return
//	}
//	fast, slow := 0, 0
//	for fast < len(nums) {
//		for slow < len(nums) && nums[slow] != 0 {
//			slow++
//		}
//		if slow > len(nums) {
//			return
//		}
//		fast = slow
//		for fast < len(nums) && nums[fast] == 0 {
//			fast++
//		}
//		if fast > len(nums) {
//			return
//		}
//		nums[slow], nums[fast] = nums[fast], nums[slow]
//	}
//}
