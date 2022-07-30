/*
LeetCode 297: https://leetcode.com/problems/serialize-and-deserialize-binary-tree/
*/

package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec struct {
}

// Constructor0297 should be Constructor. It's renamed to avoid conflicts.
func Constructor0297() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	str := fmt.Sprintf("%d", root.Val)
	str = fmt.Sprintf("%s,%s", str, this.serialize(root.Left))
	str = fmt.Sprintf("%s,%s", str, this.serialize(root.Right))
	return str
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	segs := strings.Split(data, ",")
	index := 0
	return parse(segs, &index)
}

func parse(segs []string, i *int) *TreeNode {
	seg := segs[*i]
	*i++
	if seg == "" {
		return nil
	}

	val, err := strconv.Atoi(seg)
	if err != nil {
		panic("expected error to parse an integer")
	}

	node := &TreeNode{
		Val: val,
	}
	node.Left = parse(segs, i)
	node.Right = parse(segs, i)
	return node
}
