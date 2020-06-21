/*
LeetCode 933: https://leetcode.com/problems/number-of-recent-calls/
*/

package leetcode

// RecentCounter933 is actually RecentCounter
type RecentCounter933 struct {
	queue []int
}

// Constructor933 is actually Constructor
// Rename to avoid conflicts
func Constructor933() RecentCounter933 {
	return RecentCounter933{
		queue: make([]int, 0),
	}
}

func (this *RecentCounter933) Ping(t int) int {
	this.queue = append(this.queue, t)
	i := 0
	for i < len(this.queue) && this.queue[i] < t-3000 {
		i++
	}

	this.queue = this.queue[i:]
	return len(this.queue)
}
