/*
LeetCode 427: https://leetcode.com/problems/construct-quad-tree/
*/

package leetcode

// Node0427 should be Node. It's renamed to avoid conflicts.
type Node0427 struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node0427
	TopRight    *Node0427
	BottomLeft  *Node0427
	BottomRight *Node0427
}

func construct(grid [][]int) *Node0427 {
	return helper0427(grid, 0, 0, len(grid)-1, len(grid)-1)
}

func helper0427(grid [][]int, top, left, bottom, right int) *Node0427 {
	if top == bottom {
		return &Node0427{
			Val:    grid[top][left] == 1,
			IsLeaf: true,
		}
	}

	midRow := (top + bottom + 1) / 2
	midCol := (left + right + 1) / 2
	node := &Node0427{
		IsLeaf:      false,
		TopLeft:     helper0427(grid, top, left, midRow-1, midCol-1),
		TopRight:    helper0427(grid, top, midCol, midRow-1, right),
		BottomLeft:  helper0427(grid, midRow, left, bottom, midCol-1),
		BottomRight: helper0427(grid, midRow, midCol, bottom, right),
	}

	if node.TopLeft.IsLeaf && node.TopRight.IsLeaf &&
		node.BottomLeft.IsLeaf && node.BottomRight.IsLeaf &&
		node.TopLeft.Val == node.TopRight.Val &&
		node.TopLeft.Val == node.BottomLeft.Val &&
		node.TopLeft.Val == node.BottomRight.Val {
		node = &Node0427{
			IsLeaf: true,
			Val:    node.TopLeft.Val,
		}
	}
	return node
}
