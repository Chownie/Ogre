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

	gameState := GameState{}
	gameState.Width = width
	gameState.Height = height
	gameState.Mode = "splash"
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
				if gameState.Mode == "splash" {
					gameState.Mode = "creation"
				}
			}
		}
	}
}

// CHARACTER CREATION CODE
/*character := Character{}
displaySplash(width, height, termbox.ColorCyan)
classOptions := []string{}
raceOptions := []string{}
for name, _ := range ClassBase {
	classOptions = append(classOptions, name)
}
for race, _ := range RaceBase {
	raceOptions = append(raceOptions, race)
}
raceMenuW, raceMenuH := GetMenuSize("Pick your race", raceOptions)
classMenuW, classMenuH := GetMenuSize("Pick your class", classOptions)
race := DrawMenu((width/2)-(raceMenuW/2), (height/2)-(raceMenuH/2), "Pick your race", raceOptions)
class := DrawMenu((width/2)-(classMenuW/2), (height/2)-(classMenuH/2), "Pick your class", classOptions)
character.CreateChar(class, race)
gameState.Player = &character
gameState.Mode = "mapmake"*/
