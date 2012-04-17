package main

import (
	//. "github.com/chownplusx/Termbox-Additions"
	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	width, height := termbox.Size()

	gameState := GameState{GameMap: BlankMap(40, 40), Width: width, Height: height}
	gameState.Mode = MODE_SPLASH
loop:
	for {
		termbox.Clear(termbox.ColorWhite, termbox.ColorDefault)
		gameState.GameLoop()
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break loop
			default:
				if gameState.Mode == MODE_SPLASH {
					gameState.Mode = MODE_CREATION
				}
			}
		}
	}
}
