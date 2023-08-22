package main

import "fmt"

func judgePoint24(cards []int) bool {
	list := all(cards)
	for _, v := range list {
		res := helper(v)
		for _, v := range res {
			if v == 24 {
				return true
			}
		}
	}
	return false
}
func helper(cards []int) []int {
	if len(cards) <= 1 {
		return cards
	}
	res := make([]int, 0)
	for i := 1; i < len(cards); i++ {
		left := helper(cards[:i])
		right := helper(cards[i:])

		for _, v1 := range left {
			for _, v2 := range right {
				res = append(res, v1*v2)
				if v2 != 0 {
					res = append(res, v1/v2)
				}

				res = append(res, v1+v2)
				res = append(res, v1-v2)
			}
		}
	}
	return res
}
func all(cards []int) [][]int {
	res := make([][]int, 0)
	var dfs func(target, path []int)
	dfs = func(target, path []int) {
		if len(target) == len(cards) {
			temp := make([]int, len(target))
			copy(temp, target)
			res = append(res, temp)
			return
		}
		for i := 0; i < len(path); i++ {
			newPaths := append([]int{}, path[:i]...)
			newPaths = append(newPaths, path[i+1:]...)
			dfs(append(target, path[i]), newPaths)
		}
	}
	dfs([]int{}, cards)
	return res
}
func main() {
	fmt.Println(judgePoint24([]int{4, 1, 8, 7}))
	//fmt.Println(all([]int{4, 1, 8, 7}))
}
