//package main
//
//func rotate(nums []int, k int) {
//	n := len(nums)
//	k = k % n
//	if n == 0 {
//		return
//	}
//	i, j := 0, n-1
//	for i < j {
//		nums[i], nums[j] = nums[j], nums[i]
//		i++
//		j--
//	}
//	i, j = 0, k-1
//	for i < j {
//		nums[i], nums[j] = nums[j], nums[i]
//		i++
//		j--
//	}
//	i, j = k, n-1
//	for i < j {
//		nums[i], nums[j] = nums[j], nums[i]
//		i++
//		j--
//	}
//	return
//}
