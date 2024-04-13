package utils

import "main/pkg/models"

func Contains(array []models.Position, p models.Position) bool {
	for _, i := range array {
		if i.X == p.X && i.Y == p.Y {
			return true
		}
	}
	return false
}
