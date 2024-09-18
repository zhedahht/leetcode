/*
LeetCode 478: https://leetcode.com/problems/generate-random-point-in-a-circle/description/
*/

package leetcode

import "math/rand"

// The name should be Solution. It's renamed to avoid conflicts.
type Solution478 struct {
	radius   float64
	x_center float64
	y_center float64
}

// The name should be Constructor. It's renamed to avoid conflicts.
func Constructor478(radius float64, x_center float64, y_center float64) Solution478 {
	return Solution478{
		radius:   radius,
		x_center: x_center,
		y_center: y_center,
	}
}

func (this *Solution478) RandPoint() []float64 {
	x, y := rand.Float64(), rand.Float64()
	x = x*2 - 1
	y = y*2 - 1
	if x*x+y*y > 1 {
		return this.RandPoint()
	}

	return []float64{x*this.radius + this.x_center, y*this.radius + this.y_center}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */
