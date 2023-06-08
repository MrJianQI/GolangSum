//package Code
//
//func minEatingSpeed(piles []int, h int) int {
//	res := max(piles)
//	var dfs func(piles []int, v int) int
//	dfs = func(piles []int, v int) int {
//		temp := 0
//		for _, i := range piles {
//			if i <= v {
//				temp += 1
//			} else {
//				temp += getV(i, v)
//			}
//			if temp > h {
//				return temp
//			}
//
//		}
//		return temp
//	}
//	i, j := 1, max(piles)
//	for i <= j {
//		mid := (i + j) / 2
//		temp := dfs(piles, mid)
//		if temp <= h {
//			if mid < res {
//				res = mid
//			}
//			j = mid - 1
//		} else {
//			i = mid + 1
//		}
//	}
//	return res
//}
//func max(datas []int) int {
//	res := datas[0]
//	for i := 1; i < len(datas); i++ {
//		if datas[i] > res {
//			res = datas[i]
//		}
//	}
//	return res
//}
//func getV(i, v int) int {
//	a, b := i/v, i%v
//	if b != 0 {
//		a++
//	}
//	return a
//}
