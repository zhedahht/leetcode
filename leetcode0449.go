/*
LeetCode 449: https://leetcode.com/problems/serialize-and-deserialize-bst/description/
*/

package leetcode

import (
	"strconv"
	"strings"
)

// The name should be Codec. It's renamed to avoid conflicts.
type Codec449 struct {
}

// The name should be Constructor. It's renamed to avoid conflicts.
func Constructor449() Codec449 {
	return Codec449{}
}

// Serializes a tree to a single string.
func (this *Codec449) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	node := strconv.Itoa(root.Val)
	left := this.serialize(root.Left)
	right := this.serialize(root.Right)

	for _, str := range []string{left, right} {
		if len(str) > 0 {
			node = node + "," + str
		}
	}

	return node
}

// Deserializes your encoded data to tree.
func (this *Codec449) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	strs := strings.Split(data, ",")
	nums := make([]int, len(strs))
	for i, str := range strs {
		nums[i], _ = strconv.Atoi(str)
	}

	return this.deserializeHelper(nums, 0, len(nums)-1)
}

func (this *Codec449) deserializeHelper(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	right := start + 1
	val := nums[start]
	for ; right <= end && nums[right] < val; right++ {
	}

	return &TreeNode{
		Val:   val,
		Left:  this.deserializeHelper(nums, start+1, right-1),
		Right: this.deserializeHelper(nums, right, end),
	}
}

/**
 * Your Codec449 object will be instantiated and called as such:
 * ser := Constructor449()
 * deser := Constructor449()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
