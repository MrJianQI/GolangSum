//package main
//
//type MinStack struct {
//	listN  []int
//	helper []int
//}
//
//func Constructor() MinStack {
//	return MinStack{
//		listN:  make([]int, 0),
//		helper: make([]int, 0),
//	}
//}
//
//func (this *MinStack) Push(val int) {
//	if len(this.helper) == 0 {
//		this.helper = append(this.helper, val)
//		this.listN = append(this.listN, val)
//		return
//	} else {
//		this.listN = append(this.listN, val)
//		if val <= this.helper[len(this.helper)-1] {
//			this.helper = append(this.helper, val)
//		}
//		//this.helper = append(this.helper, val)
//		return
//	}
//}
//
//func (this *MinStack) Pop() {
//	x := this.listN[len(this.listN)-1]
//	this.listN = this.listN[:len(this.listN)-1]
//	if x == this.helper[len(this.helper)-1] {
//		this.helper = this.helper[:len(this.helper)-1]
//	}
//	return
//}
//
//func (this *MinStack) Top() int {
//	return this.listN[len(this.listN)-1]
//}
//
//func (this *MinStack) GetMin() int {
//	return this.helper[len(this.helper)-1]
//}
//
//func main() {
//	obj := Constructor()
//	obj.Push(3)
//	obj.Pop()
//	param_3 := obj.Top()
//	param_4 := obj.GetMin()
//}
