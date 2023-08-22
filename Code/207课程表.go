package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	dict := map[int][]int{}
	inDegrees := make([]int, numCourses)
	for _, v := range prerequisites {
		first, after := v[0], v[1]
		dict[first] = append(dict[first], after)
		inDegrees[after]++
	}
	queue := make([]int, 0)
	num := 0
	for k, v := range inDegrees {
		if v == 0 {
			queue = append(queue, k)
		}
	}
	for len(queue) != 0 {
		num++
		v := queue[0]
		queue = queue[1:]
		for _, after := range dict[v] {
			inDegrees[after]--
			if inDegrees[after] == 0 {
				queue = append(queue, after)
			}
		}
	}
	return num == numCourses
}
func main() {
	fmt.Println(canFinish(2, [][]int{
		[]int{1, 0},
	}))
}
