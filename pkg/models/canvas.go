package models

import "fmt"

type Canvas struct {
	pixels      [][]bool
	emptyPixel  string
	filledPixel string
}

func NewCanvas(height, width int) *Canvas {
	p := make([][]bool, height)
	for i := range p {
		p[i] = make([]bool, width)
	}

	return &Canvas{
		pixels:      p,
		filledPixel: "\xF0\x9F\x9F\xA9",
		emptyPixel:  "\xF0\x9F\x9F\xAB",
	}
}

// returns height, width
func (c *Canvas) Size() (int, int) {
	return len(c.pixels), len(c.pixels[0])
}

func (c *Canvas) Clear() {
	for i := range c.pixels {
		for j := range c.pixels[i] {
			c.pixels[i][j] = false
		}
	}
}

func (c *Canvas) Update(positions []Position) error {
	for _, pos := range positions {
		c.pixels[pos.Y][pos.X] = true
	}

	return nil
}

func (c *Canvas) Draw() {
	fmt.Printf("\033c\x0c")
	for _, row := range c.pixels {
		for _, p := range row {
			switch {
			case p:
				fmt.Printf(c.filledPixel)
			default:
				fmt.Printf(c.emptyPixel)
			}
		}
		fmt.Println()
	}
}
