package main

import "fmt"

func partitionLabels(s string) []int {
	var dfs func(str string)
	res := make([]int, 0)
	dfs = func(str string) {
		if len(str) == 0 {
			return
		}
		dict := make(map[uint8]int)
		n := len(str)
		for i := n - 1; i >= 0; i-- {
			if _, ok := dict[str[i]]; !ok {
				dict[str[i]] = i
			}
		}
		i := 0
		length := dict[str[0]]
		for i <= length {
			length = max(dict[str[i]], length)
			i++
		}
		res = append(res, length+1)
		dfs(str[length+1:])
	}
	dfs(s)
	return res
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}
