package main

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0)
	i := 0
	for i < len(intervals) {
		x, y := intervals[i][0], intervals[i][1]
		j := i + 1
		for j < len(intervals) {
			if y >= intervals[j][0] {
				x = min(x, intervals[j][0])
				y = max(y, intervals[j][1])
				j++
			} else {
				break
			}

		}
		res = append(res, []int{x, y})
		i = j
	}
	return res
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
