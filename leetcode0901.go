/*
LeetCode 901: https://leetcode.com/problems/online-stock-span/
*/

package leetcode

type StockSpanner struct {
	stack [][]int
	count int
}

// Constructor901 should be Constructor.
// Rename it to avoid name conflicts.
func Constructor901() StockSpanner {
	return StockSpanner{
		stack: [][]int{{0x7fffffff, 0}},
		count: 0,
	}
}

func (this *StockSpanner) Next(price int) int {
	this.count++
	for this.stack[len(this.stack)-1][0] <= price {
		this.stack = this.stack[:len(this.stack)-1]
	}

	span := this.count - this.stack[len(this.stack)-1][1]
	this.stack = append(this.stack, []int{price, this.count})
	return span
}
