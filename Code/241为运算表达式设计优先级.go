package main

import (
	"fmt"
	"strconv"
)

func diffWaysToCompute(expression string) []int {
	if len(expression) == 0 {
		return []int{}
	}
	if isdig(expression) {
		cur, _ := strconv.Atoi(expression)
		return []int{cur}
	}

	res := make([]int, 0)
	for i, v := range expression {
		if v == '+' || v == '*' || v == '-' {
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])
			for _, v1 := range left {
				for _, v2 := range right {
					var temp int
					if expression[i] == '+' {
						temp = v1 + v2
					} else if expression[i] == '-' {
						temp = v1 - v2
					} else {
						temp = v1 * v2
					}
					res = append(res, temp)
				}
			}
		}

	}
	return res
}
func isdig(nums string) bool {
	_, err := strconv.Atoi(nums)
	if err != nil {
		return false
	}
	return true
}
func main() {
	fmt.Println(diffWaysToCompute("2-1-1"))
}
