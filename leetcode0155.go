/*
LeetCode 155: https://leetcode.com/problems/min-stack/description/
*/

package leetcode

type MinStack struct {
	stack []int
	min   []int
}

// The name should be Constructor. It's renamed to avoid conflicts.
func Constructor0155() MinStack {
	return MinStack{
		stack: make([]int, 0),
		min:   make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

	if len(this.min) == 0 {
		this.min = append(this.min, val)
	} else {
		this.min = append(this.min, min(val, this.min[len(this.min)-1]))
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
