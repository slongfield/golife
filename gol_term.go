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

// TBRenderHalf acts similar to TBRender, but uses half blocks. Doesn't use the standard render
// function, since it needs to consider two blocks at a time. Assumes a maxHP of 1.
func TBRenderHalf(l *Life) {
	for y := 0; y < l.h; y += 2 {
		for x := 0; x < l.w; x++ {
			if l.board.getHP(x, y)+l.board.getHP(x, y+1) == 2 {
				termbox.SetCell(x, y/2, '█', termbox.Attribute(0), 0)
			} else if l.board.getHP(x, y) == 1 {
				termbox.SetCell(x, y/2, '▀', termbox.Attribute(0), 0)
			} else if l.board.getHP(x, y+1) == 1 {
				termbox.SetCell(x, y/2, '▄', termbox.Attribute(0), 0)
			} else {
				termbox.SetCell(x, y/2, ' ', termbox.Attribute(1), 0)
			}
		}
	}
}
