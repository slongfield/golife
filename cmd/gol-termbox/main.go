package main

import (
	"flag"
	"time"

	"github.com/nsf/termbox-go"
	"github.com/slongfield/golife"
)

var (
	update int
)

func init() {
	flag.IntVar(&update, "update", 100, "update interval in ms")
	flag.Parse()
}

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}

	defer func() {
		termbox.Close()
		// Reset the terminal color scheme.
		termbox.SetCell(0, 0, ' ', termbox.ColorWhite, termbox.ColorBlack)
	}()

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	event := make(chan termbox.Event)

	go func() {
		for {
			event <- termbox.PollEvent()
		}
	}()

	w, h := termbox.Size()
	l := golife.NewLife(w, h, 7)
	l.Randomize()

	tick := time.Tick(time.Duration(update) * time.Millisecond)
	for {
		select {
		case <-tick:
			l.Step()
			golife.TBRainbow(l)
			termbox.Flush()
		case ev := <-event:
			switch ev.Type {
			case termbox.EventKey:
				if ev.Key == termbox.KeySpace {
					l.Randomize()
				}
				if ev.Key == termbox.KeyEsc {
					return
				}
			case termbox.EventResize:
				l.Resize(ev.Width, ev.Height)
			}
		}
	}
}
