//package main
//
//import "fmt"
//
//func lengthOfLongestSubstring(s string) int {
//	if len(s) == 0 {
//		return 0
//	} else if len(s) == 1 {
//		return 1
//	} else {
//		res := 0
//		dict := make(map[uint8]int)
//		i, j := 0, 0
//		for j < len(s) {
//			if _, ok := dict[s[j]]; !ok {
//				dict[s[j]] = 1
//				j++
//				continue
//			}
//			res = max(res, len(dict))
//			for i <= j {
//				delete(dict, s[i])
//				i++
//				if s[i-1] == s[j] {
//					break
//				}
//			}
//			dict[s[j]] = 1
//			j++
//
//		}
//		res = max(res, len(dict))
//		return res
//	}
//}
//func max(i, j int) int {
//	if i > j {
//		return i
//	}
//	return j
//}
//func main() {
//	fmt.Println(lengthOfLongestSubstring("dvdf"))
//}
