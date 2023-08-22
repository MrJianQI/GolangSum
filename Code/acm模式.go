package main

import "fmt"

//	func main() {
//		n := 5
//		list := make([]chan int, n)
//		exit := make(chan int, 1)
//		for i := 0; i < n; i++ {
//			list[i] = make(chan int, 1)
//		}
//		res := 1
//		j := 0
//		for i := 0; i < n; i++ {
//			go func(i int) {
//				for {
//					<-list[i]
//					if res > 10 {
//						exit <- 1
//						break
//					}
//					fmt.Println(i, res)
//					res++
//					if j == n-1 {
//						j = 0
//					} else {
//						j++
//					}
//					list[j] <- 1
//				}
//			}(i)
//		}
//		list[0] <- 1
//		select {
//		case <-exit:
//			fmt.Println("END")
//		}
//	}
func main() {
	n := 5
	list := make([]chan int, n)
	for i := 0; i < n; i++ {
		list[i] = make(chan int, 0)
	}
	exit := make(chan int, 1)
	res := 1
	j := 0
	for i := 0; i < n; i++ {
		go func(i int) {
			for {
				<-list[i]
				if res >= 10 {
					exit <- 1
					break
				}
				fmt.Println(i, res)
				res++
				if j == n-1 {
					j = 0
				} else {
					j++
				}
				list[j] <- 1
			}
		}(i)
	}
	list[0] <- 1
	select {
	case <-exit:
		fmt.Println("END")
	}
}
