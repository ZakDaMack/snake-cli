package main

import (
	"flag"
	"main/pkg/game"
)

func main() {
	width := flag.Int("width", 20, "Width of the game canvas")
	height := flag.Int("height", 20, "Height of the game canvas")
	speed := flag.Int("speed", 100, "Refresh rate in ms")
	flag.Parse()

	game.MakeGame(*height, *width, *speed)

}
