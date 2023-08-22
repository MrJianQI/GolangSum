//package main
//
//import "fmt"
//
//func majorityElement(nums []int) int {
//	val, count := nums[0], 1
//	for i := 1; i < len(nums); i++ {
//		fmt.Println(val, count)
//		if count == 0 {
//			val = nums[i]
//			count = 1
//			continue
//		}
//		if nums[i] == val {
//			count++
//		} else {
//			count--
//		}
//	}
//	return val
//}
//func main() {
//	fmt.Println(majorityElement([]int{10, 9, 9, 9, 10}))
//}
