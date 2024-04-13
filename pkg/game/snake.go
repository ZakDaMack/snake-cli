package game

import (
	"fmt"
	"main/pkg/models"
	"main/pkg/utils"
	"time"
)

func MakeGame(height, width, speed int) {

	// make canvas and set bounds
	c := models.NewCanvas(height, width)
	upperBounds := models.Position{X: width, Y: height}

	// make the sprites
	snake := models.NewSnake(upperBounds)
	initialPos := snake.Move()
	food := models.NewFoodItem(height, width, initialPos)

	snake.ChangeDirection(models.DIRECTION_UP)
	snake.ChangeDirection(models.DIRECTION_LEFT)

	for {
		c.Clear()              // clear canvas
		newPos := snake.Move() // move snake and add to canvas
		c.Update(newPos)

		// has the snake hit itself?, if so, quit
		if utils.Contains(newPos[:1], newPos[len(newPos)-1]) {
			c.Draw()
			fmt.Println("Game over. Score:", len(newPos))
			return
		}

		// check to see if the snake has consumed the food
		// if so, increase snakes length and spawn new food
		if utils.Contains(newPos, food) {
			snake.IncreaseLength()
			food = models.NewFoodItem(height, width, newPos)
		}

		c.Update([]models.Position{food})
		time.Sleep(time.Duration(speed) * time.Millisecond)
		c.Draw()
	}
}
