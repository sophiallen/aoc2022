package common

import "math"

type Position struct {
	Row int
	Col int
}

func (p1 Position) ManhattanDist(p2 Position) int {
	xOffset := float64(p1.Col - p2.Col)
	yOffset := float64(p1.Row - p2.Row)
	sum := math.Abs(xOffset) + math.Abs(yOffset)
	return int(sum)
}
