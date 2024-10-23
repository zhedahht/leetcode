/*
LeetCode 428: https://leetcode.com/problems/serialize-and-deserialize-n-ary-tree/description/
*/

package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

// The name should be Node. It's renamed to avoid conflicts.
type Node428 struct {
	Val      int
	Children []*Node428
}

// The name should be Codec. It's renamed to avoid conflicts.
type Codec428 struct {
}

// The name should be Constructor. It's renamed to avoid conflicts.
func Constructor428() *Codec428 {
	return &Codec428{}
}

func (this *Codec428) serialize(root *Node428) string {
	segs := make([]string, 0)
	segs = this.encode_tree(segs, root)
	return strings.Join(segs, ",")
}

func (this *Codec428) encode_tree(segs []string, node *Node428) []string {
	if node == nil {
		return segs
	}

	segs = append(segs, fmt.Sprintf("%d,%d", node.Val, len(node.Children)))
	for _, child := range node.Children {
		segs = this.encode_tree(segs, child)
	}

	return segs
}

func (this *Codec428) deserialize(data string) *Node428 {
	if data == "" {
		return nil
	}

	segs := strings.Split(data, ",")
	node := 0
	return this.decode_tree(segs, &node)
}

func (this *Codec428) decode_tree(segs []string, i *int) *Node428 {
	val, _ := strconv.Atoi(segs[*i])
	count, _ := strconv.Atoi(segs[*i+1])
	*i += 2
	node := &Node428{
		Val:      val,
		Children: make([]*Node428, 0),
	}

	for child := 0; child < count; child++ {
		node.Children = append(node.Children, this.decode_tree(segs, i))
	}

	return node
}

/**
 * Your Codec428 object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
