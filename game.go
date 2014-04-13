package main

import (
	"github.com/Chownie/Termbox-Additions"
    "github.com/Chownie/Termbox-Additions/utils"
	"github.com/nsf/termbox-go"
	"math"
	"math/rand"
	"os"
	"strconv"
)

const (
	MODE_CREATION = "creation"
	MODE_SPLASH   = "splash"
	MODE_MAPMAKE  = "mapmake"
	MODE_GAME     = "game"
)

type GameState struct {
	GameMap      *Map
	Player       *Character
	EnemyList    *Enemy
	ItemMap      []*Item
	Mode         string
	ScreenWidth  int
	ScreenHeight int
	RNG          *rand.Rand
    Messages     *[]string
}

func (gs *GameState) charCreate() {
	character := Character{}
	gs.DisplaySplash(termbox.ColorBlue)
	classOptions := []string{}
	raceOptions := []string{}

	for _, name := range ClassBase {
		classOptions = append(classOptions, name.Key)
	}
	for _, race := range RaceBase {
		raceOptions = append(raceOptions, race.Key)
	}

	class := additions.DrawMenu(gs.ScreenWidth, gs.ScreenHeight, "Pick your class", classOptions, additions.AL_CENTER, utils.CONNECT_NONE)
	race := additions.DrawMenu(gs.ScreenWidth, gs.ScreenHeight, "Pick your race", raceOptions, additions.AL_CENTER, utils.CONNECT_NONE)
	name := additions.DrawForm(gs.ScreenWidth, gs.ScreenHeight, "What's your name?", additions.AL_CENTER, utils.CONNECT_NONE, 14)
	termbox.Flush()
	character.CreateChar(class, race, name, gs)
	gs.Player = &character
}

func (gs *GameState) Controls() {
	switch ev := termbox.PollEvent(); ev.Type {
	case termbox.EventKey:
		switch ev.Ch {
		case 'h':
			gs.Player.MoveLeft(gs)
		case 'l':
			gs.Player.MoveRight(gs)
		case 'k':
			gs.Player.MoveUp(gs)
		case 'j':
			gs.Player.MoveDown(gs)
		}
		switch ev.Key {
		case termbox.KeyTab:
			_ = additions.DrawMenu(gs.ScreenWidth, gs.ScreenHeight, "Grid L: "+strconv.Itoa(len(gs.GameMap.RoomList)), []string{"W/E"}, additions.AL_CENTER, utils.CONNECT_NONE)
			termbox.Flush()
		case termbox.KeyEsc:
			choice := additions.DrawMenu(gs.ScreenWidth, gs.ScreenHeight, "Are you sure you wish to quit?", []string{"No", "Yes"}, additions.AL_CENTER, utils.CONNECT_NONE)
			if choice == 1 {
				termbox.Close()
				os.Exit(0)
			}
		default:
		}
	}
}

func (gs *GameState) LightArea() {
	for i := 0; i < 360; i += 4 {
		radians := (math.Pi * float64(i)) / 180
		for l := 0; l < gs.Player.Per; l++ {
			dx := math.Cos(radians) * float64(l)
			dy := math.Sin(radians) * float64(l)
			x := int(dx)
			y := int(dy)
			tile, exists := gs.GameMap.LocateTile(gs.Player.X+x, gs.Player.Y+y)
			if exists == true {
				if tile.IsWall == false {
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
		gs.DisplaySplash(termbox.ColorRed)
		termbox.Flush()
		// No need for a mode change here, it's done elsewhere

	case MODE_MAPMAKE:
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		gs.charCreate()
		gs.GenMap()
		termbox.Flush()
		gs.Mode = MODE_CREATION

	case MODE_CREATION:
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		gs.LightArea()
		termbox.Flush()
		gs.Mode = MODE_GAME

	case MODE_GAME:
		termbox.Clear(termbox.ColorDefault, termbox.ColorBlack)
        gs.DisplayStatus()
		gs.LightArea()
		gs.DisplayMap()
        gs.DisplayInventory()
        gs.DisplayMessages()
		gs.Controls()
		gs.Player.DisplayPlayer()
		termbox.Flush()

	default:
		panic(gs.Mode)
	}
}
