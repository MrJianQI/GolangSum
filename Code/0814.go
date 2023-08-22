package main

import "fmt"

func main() {
	nums := []int{10, 5, 2, -2, -13, -3}

	stack := make([]int, 0)
	for _, v := range nums {
		if len(stack) == 0 {
			stack = append(stack, v)
			continue
		}
		temp := stack[len(stack)-1]
		if temp*v > 0 {
			stack = append(stack, v)
		} else {
			if v >= 0 && temp < 0 {
				stack = append(stack, v)
			}
			if abs(temp) > abs(v) {
				continue
			} else if abs(temp) == abs(v) {
				stack = stack[:len(stack)-1]
			} else {
				for len(stack) > 0 && abs(stack[len(stack)-1]) < abs(v) {
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, v)
			}
		}
	}
	fmt.Println(stack)
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
