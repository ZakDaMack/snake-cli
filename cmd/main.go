package main

import (
	"flag"
	"main/pkg/game"
)

func main() {
	width := flag.Int("width", 20, "")
	height := flag.Int("height", 20, "")
	speed := flag.Int("speed", 100, "")
	flag.Parse()

	game.MakeGame(*height, *width, *speed)

}
