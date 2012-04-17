package main

import (
	. "github.com/chownplusx/Termbox-Additions"
	"github.com/nsf/termbox-go"
)

const (
	MODE_CREATION = "creation"
	MODE_SPLASH   = "splash"
	MODE_MAPMAKE  = "mapmake"
	MODE_GAME     = "game"
)

type GameState struct {
	GameMap   *Map
	Player    *Character
	EnemyList *Enemy
	Mode      string
	Width     int
	Height    int
}

func (gs *GameState) charCreate() {
	character := Character{}
	displaySplash(gs.Width, gs.Height, termbox.ColorCyan)
	classOptions := []string{}
	raceOptions := []string{}

	for _, name := range ClassBase {
		classOptions = append(classOptions, name.Key)
	}
	for _, race := range RaceBase {
		raceOptions = append(raceOptions, race.Key)
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

func (gs *GameState) Controls() {
	switch ev := termbox.PollEvent(); ev.Type {
	case termbox.EventKey:
		switch ev.Key {
		case termbox.KeyArrowLeft:
			gs.Player.MoveLeft(gs)
		case termbox.KeyArrowRight:
			gs.Player.MoveRight(gs)
		case termbox.KeyArrowUp:
			gs.Player.MoveUp(gs)
		case termbox.KeyArrowDown:
			gs.Player.MoveDown(gs)
		}
	}
}

func (gs *GameState) GameLoop() {
	switch gs.Mode {
	case MODE_SPLASH:
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		displaySplash(gs.Width, gs.Height, termbox.ColorRed)
		termbox.Flush()
		// No need for a mode change here, it's done elsewhere
	case MODE_CREATION:
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		gs.charCreate()
		termbox.Flush()
		gs.Mode = MODE_MAPMAKE

	case MODE_MAPMAKE:
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		gs.GameMap.MakeRoom(10, false)
		termbox.Flush()
		gs.Mode = MODE_GAME

	case MODE_GAME:
		gs.DisplayMap()
		gs.Controls()
		gs.Player.DisplayPlayer()
		termbox.Flush()

	default:
		panic(gs.Mode)
	}
}
