package main

import (
	additions "github.com/Chownie/Termbox-Additions"
	"github.com/Chownie/Termbox-Additions/utils"
    "github.com/nsf/termbox-go"
	"log"
	"math/rand"
	"os"
	"time"
)

var (
	Loggy *log.Logger
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	width, height := termbox.Size()

	f, err := os.OpenFile("errLog.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("error opening/writing file")
	}
	defer f.Close()

	Loggy = log.New(f, "", 0)

	rng := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	gameState := new(GameState)
    gameState.GameMap = BlankMap(80, 20)
	gameState.ScreenWidth = width
	gameState.ScreenHeight = height
	gameState.RNG = rng
    gameState.Messages = &[]string{"", "", ""}
	gameState.Mode = MODE_SPLASH

	for {
		termbox.Clear(termbox.ColorWhite, termbox.ColorDefault)
		gameState.GameLoop()
		if gameState.Mode == MODE_SPLASH {
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeyEsc:
					choice := additions.DrawMenu(width, height, "Are you sure you want to quit?", []string{"Yes", "No"}, additions.AL_CENTER, utils.CONNECT_NONE)
					if choice == 0 {
						termbox.Close()
						os.Exit(0)
					}
				default:
					gameState.Mode = MODE_MAPMAKE
				}
			}
		}
	}
}
