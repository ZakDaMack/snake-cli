package game

import (
	"main/internal/utils"
	"main/pkg/models"
	"time"
)

func MakeGame(height, width, speed int) {

	// make canvas and set bounds
	c := models.NewCanvas(height, width)
	upperBounds := models.Position{X: width, Y: height}

	// make the sprites
	snake := models.NewSnake(upperBounds)
	// initialPos := snake.Move()
	// food := models.NewFoodItem(height, width, initialPos)
	food := models.NewFixedItem(width-2, height/2)

	for {
		c.Clear()
		newPos := snake.Move()
		c.Update(newPos)
		if utils.Contains(newPos, food) {
			snake.IncreaseLength()
			food = models.NewFoodItem(height, width, newPos)
		}
		c.Update([]models.Position{food})
		time.Sleep(time.Duration(speed) * time.Millisecond)
		c.Draw()
	}
}
