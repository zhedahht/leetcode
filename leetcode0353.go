/*
LeetCode 353: https://leetcode.com/problems/design-snake-game/
*/

package leetcode

import "strings"

type SnakeGame struct {
	width     int
	height    int
	board     [][]bool
	snake     [][]int
	pos       []int
	food      [][]int
	foodIndex int
}

/** Initialize your data structure here.
  @param width - screen width
  @param height - screen height
  @param food - A list of food positions
  E.g food = [[1,1], [1,0]] means the first food is positioned at [1,1], the second is at [1,0]. */
func Constructor(width int, height int, food [][]int) SnakeGame {
	game := SnakeGame{}
	game.width = width
	game.height = height
	game.board = make([][]bool, height)
	for i := 0; i < height; i++ {
		game.board[i] = make([]bool, width)
	}
	game.board[0][0] = true

	game.pos = []int{0, 0}
	game.snake = make([][]int, 0)
	game.snake = append(game.snake, []int{0, 0})
	game.food = food
	game.foodIndex = 0

	return game
}

/** Moves the snake.
  @param direction - 'U' = Up, 'L' = Left, 'R' = Right, 'D' = Down
  @return The game's score after the move. Return -1 if game over.
  Game over when snake crosses the screen boundary or bites its body. */
func (this *SnakeGame) Move(direction string) int {
	dirs := [][]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
	directions := "ULRD"
	dir := dirs[strings.Index(directions, direction)]

	this.pos = []int{this.pos[0] + dir[0], this.pos[1] + dir[1]}
	if this.pos[0] < 0 || this.pos[0] >= this.height || this.pos[1] < 0 || this.pos[1] >= this.width {
		return -1
	}

	if this.foodIndex >= len(this.food) || this.pos[0] != this.food[this.foodIndex][0] || this.pos[1] != this.food[this.foodIndex][1] {
		tail := this.snake[0]
		this.snake = this.snake[1:]
		this.board[tail[0]][tail[1]] = false
	} else {
		this.foodIndex++
	}

	if this.board[this.pos[0]][this.pos[1]] {
		return -1
	}

	this.board[this.pos[0]][this.pos[1]] = true
	this.snake = append(this.snake, this.pos)

	return this.foodIndex
}

/**
 * Your SnakeGame object will be instantiated and called as such:
 * obj := Constructor(width, height, food);
 * param_1 := obj.Move(direction);
 */
