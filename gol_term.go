package golife

import (
	termbox "github.com/nsf/termbox-go"
)

// This file contains the termbox rendering code for a game of life cell.

// TBRender is a rendering function that renders the board in using termbox.
// Doesn't make any assumptions about the colors, and so it can look pretty
// crazy. =D
func TBRender(l *Life) {
	l.Render(func(x, y, hp int) {
		termbox.SetCell(x, y, '█', termbox.Attribute(hp), 0)
	})
}

// TBRainbow attempts to draw a rainbow. Assumes that the maxHP is 7
func TBRainbow(l *Life) {
	colors := []termbox.Attribute{
		termbox.ColorBlack,
		termbox.ColorRed,
		termbox.ColorGreen | termbox.AttrReverse,
		termbox.ColorYellow,
		termbox.ColorGreen,
		termbox.ColorBlue,
		termbox.ColorCyan,
		termbox.ColorWhite,
	}
	l.Render(func(x, y, hp int) {
		termbox.SetCell(x, y, '█', colors[hp], 0)
	})
}
