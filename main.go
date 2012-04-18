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

	gameState := GameState{GameMap: BlankMap(50, 50), Width: width, Height: height}
	gameState.Mode = MODE_SPLASH
loop:
	for {
		termbox.Clear(termbox.ColorWhite, termbox.ColorDefault)
		gameState.GameLoop()
		if gameState.Mode == MODE_SPLASH {
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeyEsc:
					break loop
				default:
					gameState.Mode = MODE_CREATION
				}
			}
		}
	}
}
