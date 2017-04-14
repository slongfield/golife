// gol-text is a simple text rendering for the Game of Life library.
package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/slongfield/golife"
)

var (
	width  int
	height int
	update int
)

func init() {
	flag.IntVar(&width, "width", 100, "width of the board")
	flag.IntVar(&height, "height", 70, "height of the board")
	flag.IntVar(&update, "update", 100, "update interval in ms")
	flag.Parse()
}

func main() {
	tick := time.Tick(time.Duration(update) * time.Millisecond)

	l := golife.NewLife(width, height, 4)
	l.SetTextSprite([]string{" ", "░", "▒", "▓", "█"})
	l.Randomize()

	// Clear the screen, and hide the cursor. Not cross-platform.
	fmt.Print("\033[2J\033[?25l")

	for _ = range tick {
		fmt.Print("\033[H", l) // Clear screen and print. Not cross-platform.
		l.Step()
	}
}
