/*
LeetCode 641: https://leetcode.com/problems/design-circular-deque/description/
*/

package leetcode

type MyCircularDeque struct {
	data       []int
	frontIndex int
	lastIndex  int
	length     int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		data:       make([]int, k),
		frontIndex: 0,
		lastIndex:  k - 1,
		length:     0,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}

	this.data[this.frontIndex] = value
	this.frontIndex = (this.frontIndex + 1) % len(this.data)
	this.length++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}

	this.data[this.lastIndex] = value
	this.lastIndex = (this.lastIndex - 1 + len(this.data)) % len(this.data)
	this.length++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}

	this.frontIndex = (this.frontIndex - 1 + len(this.data)) % len(this.data)
	this.length--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}

	this.lastIndex = (this.lastIndex + 1) % len(this.data)
	this.length--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.length == 0 {
		return -1
	}

	index := (this.frontIndex - 1 + len(this.data)) % len(this.data)
	return this.data[index]
}

func (this *MyCircularDeque) GetRear() int {
	if this.length == 0 {
		return -1
	}

	index := (this.lastIndex + 1) % len(this.data)
	return this.data[index]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.length == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.length == len(this.data)
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
