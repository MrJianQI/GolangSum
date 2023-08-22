package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(candidates)
	var dfs func(path []int, flag int)
	dfs = func(path []int, flag int) {
		total := sum(path)
		if total == target {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for _, v := range candidates {
			if v < flag {
				continue
			}
			if total+v > target {
				break
			}
			dfs(append(path, v), v)
		}
	}
	dfs([]int{}, candidates[0])
	return res
}
func sum(nums []int) int {
	res := 0
	for _, v := range nums {
		res += v
	}
	return res
}
func main() {
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}
