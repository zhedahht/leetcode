/*
LeetCode 427: https://leetcode.com/problems/construct-quad-tree/
*/

package leetcode

// The type name should be Node. It's renamed to avoid conflicts.
type Node0429 struct {
	Val      int
	Children []*Node0429
}

// The function name should be levelOrder. It's renamed to avoid conflicts.
func levelOrder0429(root *Node0429) [][]int {
	result := make([][]int, 0)
	queue1, queue2 := make([]*Node0429, 0), make([]*Node0429, 0)
	if root != nil {
		queue1 = append(queue1, root)
		result = append(result, make([]int, 0))
	}

	for len(queue1) > 0 {
		node := queue1[0]
		queue1 = queue1[1:]
		result[len(result)-1] = append(result[len(result)-1], node.Val)

		for _, child := range node.Children {
			queue2 = append(queue2, child)
		}

		if len(queue1) == 0 {
			queue1 = queue2
			queue2 = make([]*Node0429, 0)
			if len(queue1) > 0 {
				result = append(result, make([]int, 0))
			}
		}
	}

	return result
}
