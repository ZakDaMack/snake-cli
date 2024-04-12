package models

import "math/rand"

func NewFoodItem(maxHeight int, maxWidth int, exclusionPosition []Position) Position {
	return Position{
		X: rand.Intn(maxWidth),
		Y: rand.Intn(maxHeight),
	}
}

func NewFixedItem(x, y int) Position {
	return Position{
		X: x,
		Y: y,
	}
}
