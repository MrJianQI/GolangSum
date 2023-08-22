//package main
//
//func isValid(s string) bool {
//	if len(s)%2 == 1 {
//		return false
//	}
//	list := make([]int32, 0)
//	for _, v := range s {
//		if v == '(' || v == '[' || v == '{' {
//			list = append(list, v)
//		} else if v == ')' && (len(list) == 0 || list[len(list)-1] != '(') {
//			return false
//		} else if v == ']' && (len(list) == 0 || list[len(list)-1] != '[') {
//			return false
//		} else if v == '}' && (len(list) == 0 || list[len(list)-1] != '{') {
//			return false
//		} else {
//			list = list[:len(list)-1]
//		}
//	}
//	if len(list) != 0 {
//		return false
//	}
//	return true
//}
