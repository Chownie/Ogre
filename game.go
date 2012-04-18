package main

import (
	. "github.com/chownplusx/Termbox-Additions"
	"github.com/nsf/termbox-go"
	"math"
	"os"
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
		case termbox.KeyEsc:
			os.Exit(0)
		default:
			print(ev.Key)
			print("\n")
			print(ev.Ch)
			print("\n")
			print("\n")
		}
	}
}

func (gs *GameState) LightArea() {
	for i := 0; i < 360; i++ {
		radians := (math.Pi * float64(i)) / 180
		for l := 0; l < gs.Player.Per; l++ {
			dx := math.Cos(radians) * float64(l)
			dy := math.Sin(radians) * float64(l)
			x := int(dx)
			y := int(dy)
			tile, exists := gs.GameMap.LocateTile(gs.Player.X+x, gs.Player.Y+y)
			if exists == true {
				if tile.IsWalkable {
					tile.IsVisible = true
				} else {
					tile.IsVisible = true
					break
				}
			} else {
				break
			}
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
		gs.GameMap.MakeRoom(0, false)
		termbox.Flush()
		gs.Mode = MODE_GAME

	case MODE_GAME:
		termbox.Clear(termbox.ColorDefault, termbox.ColorBlack)
		gs.LightArea()
		gs.DisplayMap()
		gs.Controls()
		gs.Player.DisplayPlayer()
		termbox.Flush()

	default:
		panic(gs.Mode)
	}
}
