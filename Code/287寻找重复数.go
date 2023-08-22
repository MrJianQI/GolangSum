//package main
//
//import "fmt"
//
//func findDuplicate(nums []int) int {
//	slow := nums[0]
//	fast := nums[nums[0]]
//
//	for slow != fast {
//		slow = nums[slow]
//		fast = nums[nums[fast]]
//	}
//	slow = 0
//	for slow != fast {
//		slow = nums[slow]
//		fast = nums[fast]
//	}
//	return slow
//}
//func main() {
//	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
//
//}
