package main

import (
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
)

func main() {
	// width := flag.Int("width", 20, "")
	// height := flag.Int("height", 20, "")
	// speed := flag.Int("speed", 100, "")
	// flag.Parse()

	// game.MakeGame(*height, *width, *speed)

	keyEvent, err := keyboard.GetKeys(1)
	if err != nil {
		fmt.Println("error reading input:", err)
		os.Exit(1)
	}
	for {
		event := <-keyEvent
		if event.Err != nil {
			panic(event.Err)
		}
		fmt.Printf("You pressed: rune %q, key %X\r\n", event.Rune, event.Key)
		// if event.Key == keyboard.KeyArrowRight
		// if event.Key == keyboard.KeyArrowLeft
		// if event.Key == keyboard.KeyArrowDown
		// if event.Key == keyboard.KeyArrowUp

		if event.Key == keyboard.KeyEsc {
			break
		}
	}
}
