/*
LeetCode 622: https://leetcode.com/problems/design-circular-queue/
*/

package leetcode

type MyCircularQueue struct {
	queue      []int
	frontIndex int
	rearIndex  int
	empty      bool
	full       bool
}

/** Initialize your data structure here. Set the size of the queue to be k. */
// The function name below should be Constructor. Rename it to avoid name conflict with leetcode0353.
func Constructor0622(k int) MyCircularQueue {
	circularQueue := MyCircularQueue{}
	circularQueue.queue = make([]int, k)
	circularQueue.frontIndex = 0
	circularQueue.rearIndex = -1
	circularQueue.empty = false
	circularQueue.full = false
	return circularQueue
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	this.rearIndex = (this.rearIndex + 1) % len(this.queue)
	this.queue[this.rearIndex] = value

	this.empty = false
	if (this.rearIndex+1)%len(this.queue) == this.frontIndex {
		this.full = true
	}

	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	this.frontIndex = (this.frontIndex + 1) % len(this.queue)

	this.full = false
	if (this.rearIndex+1)%len(this.queue) == this.frontIndex {
		this.empty = true
	}

	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}

	return this.queue[this.frontIndex]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.queue[this.rearIndex]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.empty
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.full
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
