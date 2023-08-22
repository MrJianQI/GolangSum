//package main
//
//import "fmt"
//
//type Trie struct {
//	Val   rune
//	IsEnd bool
//	Data  [26]*Trie
//}
//
//func Constructor() Trie {
//	res := Trie{
//		Val:   0,
//		IsEnd: false,
//		Data:  [26]*Trie{},
//	}
//	return res
//}
//
//func (this *Trie) Insert(word string) {
//	if this.Search(word) {
//		return
//	} else {
//		n := len(word)
//		cur := this
//		for k, v := range word {
//			if cur.Data[v-'a'] == nil {
//				cur.Data[v-'a'] = &Trie{
//					Val:   v,
//					IsEnd: false,
//					Data:  [26]*Trie{},
//				}
//			}
//			cur = cur.Data[v-'a']
//			if k == n-1 {
//				cur.IsEnd = true
//			}
//		}
//		return
//	}
//}
//
//func (this *Trie) Search(word string) bool {
//	cur := this
//	for _, v := range word {
//		if cur.Data[v-'a'] == nil {
//			return false
//		}
//		cur = cur.Data[v-'a']
//	}
//	return cur.IsEnd
//}
//
//func (this *Trie) StartsWith(prefix string) bool {
//	cur := this
//	for _, v := range prefix {
//		if cur.Data[v-'a'] == nil {
//			return false
//		}
//		cur = cur.Data[v-'a']
//	}
//	return true
//}
//func main() {
//	res := Constructor()
//	res.Insert("apple")
//	fmt.Println(res.Search("apple"))
//	fmt.Println(res.Search("app"))
//	fmt.Println(res.StartsWith("app"))
//}
