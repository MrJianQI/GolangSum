//package main
//
//import "fmt"
//
//func main() {
//	n := 0
//	fmt.Scan(&n)
//	nums := make([]int, n)
//	dict := make(map[int]int)
//	for i := 0; i < n; i++ {
//		n, _ := fmt.Scan(&nums[i])
//		if n == 0 {
//			break
//		}
//		dict[nums[i]] = i
//	}
//	var x, y int
//	_, _ = fmt.Scan(&x, &y)
//	v1, ok1 := dict[x]
//	v2, ok2 := dict[y]
//	if ok1 == false || ok2 == false {
//		fmt.Println("No")
//		return
//	} else {
//		if v1+1 == v2 || v2+1 == v1 {
//			fmt.Println("Yes")
//			return
//		}
//		fmt.Println("No")
//		return
//	}
//}
