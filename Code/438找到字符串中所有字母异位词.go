//package main
//
//import "fmt"
//
//func findAnagrams(s string, p string) []int {
//	if len(s) < len(p) {
//		return []int{}
//	}
//	res := []int{}
//	i := 0
//
//	num := [26]int{}
//	dict := make(map[int32]int)
//	for _, v := range p {
//		num[v-'a'] += 1
//		dict[v] = 1
//	}
//	for i < len(s)-len(p) {
//		if _, ok := dict[int32(s[i])]; !ok {
//			i++
//			continue
//		}
//		if getOK(s[i:i+len(p)], &num) {
//			res = append(res, i)
//		}
//		i++
//	}
//	return res
//}
//func getOK(s string, dict *[26]int) bool {
//	num := [26]int{}
//	for _, v := range s {
//		num[v-'a'] += 1
//	}
//	return num == *dict
//}
//func main() {
//	fmt.Println(findAnagrams("cbaebabacd", "abc"))
//}
