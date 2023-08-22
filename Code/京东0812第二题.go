package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAX_N = 200000 + 1
const INF = 100000000

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var nums []int
	var dp [MAX_N][10]int
	n := 0
	if scanner.Scan() {
		inputs := strings.Split(scanner.Text(), " ")
		n, _ = strconv.Atoi(inputs[0])
		if scanner.Scan() {
			inputNums := strings.Split(scanner.Text(), " ")
			for _, v := range inputNums {
				temp, _ := strconv.Atoi(v)
				nums = append(nums, temp%10)
			}
		}
	}

	//dp[n-1][nums[n-1]] = 1
	if n == 1 {
		for i := 0; i < 10; i++ {
			fmt.Print(dp[0][i])
			if i < 9 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
		return
	}
	var dfs func(target []int)
	dfs = func(target []int) {
		if len(target) == 0 {
			return
		}
		if len(target) == 1 {
			dp[0][target[0]]++
			dp[0][target[0]] %= 10
			return
		}
		length := len(target)
		temp := make([]int, length-2)
		copy(temp, target[:length-2])
		a, b := target[length-1], target[length-2]
		dp[length-1][a] += 1
		dfs(append(temp, (a*b)%10))
		dfs(append(temp, (a+b)%10))
	}
	dfs(nums)
	//for i := n - 1; i >= 0; i-- {
	//	fmt.Println(dp[i])
	//}
	for i := 0; i < 10; i++ {
		fmt.Print(dp[0][i])
		if i < 9 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
