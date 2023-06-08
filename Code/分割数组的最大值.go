//package Code
//
//import "fmt"
//
//func splitArray(nums []int, k int) int {
//	i := max(nums)
//	j := sum(nums)
//	res := j
//	for i <= j {
//		mid := (i + j) / 2
//		split := split_num(nums, mid)
//		if split <= k {
//			if mid < res {
//				res = mid
//			}
//			j = mid - 1
//		} else if split > k {
//			i = mid + 1
//		}
//	}
//	return res
//}
//func split_num(nums []int, total int) int {
//	if len(nums) < 1 {
//		return 0
//	}
//	k := 0
//	res := 1
//	for _, v := range nums {
//		if k+v <= total {
//			k += v
//		} else {
//			k = v
//			res++
//		}
//	}
//	return res
//}
//func sum(nums []int) int {
//	res := 0
//	for _, v := range nums {
//		res += v
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
//	fmt.Println(splitArray([]int{1, 2, 3, 4, 5}, 2))
//}
