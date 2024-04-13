package game

import (
	"fmt"
	"main/pkg/models"
	"main/pkg/utils"
	"os"
	"time"

	"github.com/eiannone/keyboard"
)

func MakeGame(height, width, speed int) {
	// make canvas and set bounds
	c := models.NewCanvas(height, width)
	upperBounds := models.Position{X: width, Y: height}

	// make the sprite
	snake := models.NewSnake(upperBounds)

	go run(speed, snake, c)

	keyEvent, err := keyboard.GetKeys(10)
	if err != nil {
		fmt.Println("error reading input:", err)
		os.Exit(1)
	}

	defer keyboard.Close()

	for {
		event := <-keyEvent
		if event.Err != nil {
			panic(event.Err)
		}

		switch {
		case event.Key == keyboard.KeyArrowRight, string(event.Rune) == "d":
			snake.ChangeDirection(models.DIRECTION_RIGHT)
		case event.Key == keyboard.KeyArrowLeft, string(event.Rune) == "a":
			snake.ChangeDirection(models.DIRECTION_LEFT)
		case event.Key == keyboard.KeyArrowUp, string(event.Rune) == "w":
			snake.ChangeDirection(models.DIRECTION_UP)
		case event.Key == keyboard.KeyArrowDown, string(event.Rune) == "s":
			snake.ChangeDirection(models.DIRECTION_DOWN)
		}

		if event.Key == keyboard.KeyEsc {
			break
		}
	}
}

func run(speed int, snake *models.Snake, canvas *models.Canvas) {
	height, width := canvas.Size()
	initialPos := snake.Move()
	food := models.NewFoodItem(height, width, initialPos)

	for {
		canvas.Clear()         // clear canvas
		newPos := snake.Move() // move snake and add to canvas
		canvas.Update(newPos)

		// has the snake hit itself?, if so, quit
		// newPos[:len(newPos)-1] get all up to last item
		if utils.Contains(newPos[:len(newPos)-1], newPos[len(newPos)-1]) {
			canvas.Draw()
			fmt.Println("Game over. Score:", len(newPos), "(Press ESC to exit)")
			return
			// os.Exit(0)
		}

		// check to see if the snake has consumed the food
		// if so, increase snakes length and spawn new food
		if utils.Contains(newPos, food) {
			snake.IncreaseLength()
			food = models.NewFoodItem(height, width, newPos)
		}

		canvas.Update([]models.Position{food})
		time.Sleep(time.Duration(speed) * time.Millisecond)
		canvas.Draw()
	}
}
