package main

import "fmt"

func searchRange(nums []int, target int) []int {
	i, j := findLeft(nums, target), findRight(nums, target)
	return []int{i, j}
}
func findLeft(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == target {
			if mid == 0 || (nums[mid-1] != nums[mid]) {
				return mid
			} else {
				j = mid - 1
			}
		} else if nums[mid] > target {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return -1
}
func findRight(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == target {
			if mid == len(nums)-1 || (nums[mid+1] != nums[mid]) {
				return mid
			} else {
				i = mid + 1
			}
		} else if nums[mid] > target {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return -1
}
func main() {
	fmt.Println(find([]int{5, 7, 7, 8, 8, 8, 8, 8, 10}, 8))
}
