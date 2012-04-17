package main

import (
	. "github.com/chownplusx/Termbox-Additions"
	"github.com/nsf/termbox-go"
)

type GameState struct {
	GameMap   *Map
	Player    *Character
	EnemyList *Enemy
	Mode      string
	Width     int
	Height    int
}

func (gs GameState) charCreate() {
	character := Character{}
	displaySplash(gs.Width, gs.Height, termbox.ColorCyan)
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

	race := DrawMenu((gs.Width/2)-(raceMenuW/2),
		(gs.Height/2)-(raceMenuH/2), "Pick your race", raceOptions)
	class := DrawMenu((gs.Width/2)-(classMenuW/2),
		(gs.Height/2)-(classMenuH/2), "Pick your class", classOptions)
	termbox.Flush()
	character.CreateChar(class, race)
	gs.Player = &character
}

func (gs GameState) Controls() {
	switch ev := termbox.PollEvent(); ev.Type {
	case termbox.EventKey:
		switch ev.Key {
		case termbox.KeyArrowLeft:
			gs.Player.MoveLeft(&gs)
		case termbox.KeyArrowRight:
			gs.Player.MoveRight(&gs)
		case termbox.KeyArrowUp:
			gs.Player.MoveUp(&gs)
		case termbox.KeyArrowDown:
			gs.Player.MoveDown(&gs)
		}
	}
}

func (gs GameState) GameLoop() {
	switch gs.Mode {
	case "splash":
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		displaySplash(gs.Width, gs.Height, termbox.ColorRed)
		termbox.Flush()

	case "creation":
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		gs.charCreate()
		gs.Mode = "mapmake"

	case "mapmake":
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		gs.GameMap.MakeRoom(10, false)
		termbox.Flush()
		gs.Mode = "game"

	case "game":
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		gs.Controls()
		gs.Player.DisplayPlayer()
		termbox.Flush()

	default:
		panic(gs.Mode)
	}
}
