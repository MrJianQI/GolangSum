//package main
//
//import "fmt"
//
//type Node struct {
//	Val    int
//	Next   *Node
//	Random *Node
//}
//
//func copyRandomList(head *Node) *Node {
//	oldDict := make(map[*Node]int)
//	newDict := make(map[int]*Node)
//
//	newPre := &Node{}
//	oldCur, newCur := head, newPre
//	index := 0
//	for oldCur != nil {
//		oldDict[oldCur] = index
//		temp := &Node{Val: oldCur.Val}
//		newDict[index] = temp
//		index++
//		oldCur = oldCur.Next
//		newCur.Next = temp
//		newCur = newCur.Next
//	}
//	oldCur, newCur = head, newPre.Next
//	for oldCur != nil {
//		if oldCur.Random != nil {
//
//			index = oldDict[oldCur.Random]
//			newCur.Random = newDict[index]
//			fmt.Println(newCur.Val, newDict[index].Val)
//		}
//
//		oldCur = oldCur.Next
//		newCur = newCur.Next
//	}
//	return newPre.Next
//}
