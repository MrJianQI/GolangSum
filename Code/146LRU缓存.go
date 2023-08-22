package main

type LinkNode struct {
	Key  int
	Val  int
	Prev *LinkNode
	Next *LinkNode
}
type LRUCache struct {
	Size int
	Cap  int
	Head *LinkNode
	Tail *LinkNode
	Hash map[int]*LinkNode
}

func Constructor(capacity int) LRUCache {
	res := LRUCache{
		Size: 0,
		Cap:  capacity,
		Head: &LinkNode{},
		Tail: &LinkNode{},
		Hash: make(map[int]*LinkNode),
	}
	res.Head.Prev = res.Tail
	res.Head.Next = res.Tail
	res.Tail.Prev = res.Head
	res.Tail.Next = res.Head
	return res
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.Hash[key]
	if ok {
		this.MoveFirst(v)
		return v.Val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	v, ok := this.Hash[key]
	if ok {
		v.Val = value
		this.MoveFirst(v)
	} else {
		temp := &LinkNode{Val: value, Key: key}
		if this.Size < this.Cap {
			this.InsertNode(temp)
			this.Hash[key] = temp
		} else {
			this.DeleteNode()
			this.InsertNode(temp)
			this.Hash[key] = temp
		}
	}
}
func (this *LRUCache) MoveFirst(v *LinkNode) {
	prev, nxt := v.Prev, v.Next

	prev.Next = nxt
	nxt.Prev = prev

	nxt = this.Head.Next

	v.Next = nxt
	nxt.Prev = v

	this.Head.Next = v
	v.Prev = this.Head
}
func (this *LRUCache) DeleteNode() {
	this.Size -= 1
	v := this.Tail.Prev
	delete(this.Hash, v.Key)
	//for k, _ := range this.Hash {
	//	fmt.Println(k)
	//}
	prev, nxt := v.Prev, v.Next
	prev.Next = nxt
	nxt.Prev = prev
}
func (this *LRUCache) InsertNode(v *LinkNode) {
	this.Size += 1
	prev, nxt := this.Head, this.Head.Next
	v.Prev = prev
	v.Next = nxt

	prev.Next = v
	nxt.Prev = v
}
