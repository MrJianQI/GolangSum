//package main
//
//func groupAnagrams(strs []string) [][]string {
//	res := make([][]string, 0)
//	if len(strs) == 0 {
//		res = append(res, []string{""})
//		return res
//	} else if len(strs) == 1 {
//		res = append(res, []string{strs[0]})
//		return res
//	} else {
//		for i := 0; i < len(strs); i++ {
//			if len(res) == 0 {
//				res = append(res, []string{strs[i]})
//				continue
//			} else {
//				for j, v := range res {
//					if isOK(v[0], strs[i]) {
//						res[j] = append(res[j], strs[i])
//						continue
//					}
//				}
//				res = append(res, []string{strs[i]})
//			}
//		}
//		return res
//	}
//}
//
//func isOK(old, new string) bool {
//	if len(old) != len(new) {
//		return false
//	}
//	oldArray := make([]int, 26)
//	newArray := make([]int, 26)
//	for i := 0; i < len(old); i++ {
//		oldArray[old[i]-'a'] += 1
//		newArray[new[i]-'a'] += 1
//	}
//	for i := 0; i < 26; i++ {
//		if oldArray[i] != newArray[i] {
//			return false
//		}
//	}
//	return true
//}
