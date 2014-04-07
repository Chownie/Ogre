package main

import (
	//. "github.com/chownplusx/Termbox-Additions"
	. "github.com/Chownie/Termbox-Additions/utils"
	"github.com/nsf/termbox-go"
	"strconv"
	"strings"

//	"fmt"
)

const (
	SPLASH = "\n\n\n\n\n\n @@@@@@@@    @@@@@@@@  @@@@@@@@@   @@@@@@@@  \n@@@@@@@@@@  @@@@@@@@@  @@@@@@@@@@  @@@@@@@@  \n@@!    @@@  !@@        @@!    @@@  @@!       \n!@!    @!@  !@!        !@!    @!@  !@!       \n@!@    !@!  !@! @!@!@  @!@!!!!@!   @@!!!:!   \n!@!    !!!  !@! !!@!!  !!@@@!@!    !!!!!!:   \n!!:    !!!  !!!   !!:  !!:: :!!!   !!:       \n:!:    !:!  :!!   !::  !!:    !:!  :!:       \n!!::::!:!   ::! :!!:   :!     ::   :: ::::   \n:!::  :!:   ;: :: :    :      :    :   ;::   \n:::    :;    :    :                '   '     \n ::          '                               \n  '\n\n\n\n\n"
)

func displaySplash(width, height int, fg termbox.Attribute) {
	x := (width / 2) - (43 / 2)
	y := (height / 2) - (13 / 1)
	DrawRichTextMulti(x, y, SPLASH, fg, termbox.ColorDefault)
	DrawText(x+12, y+24, "Press any key to start")
}

func (ch *Character) DisplayPlayer() {
	DrawRichText(ch.X, ch.Y, "@", termbox.ColorGreen, termbox.ColorBlack)
}

func (gs *GameState) DisplayStatus() {
	hp := strconv.Itoa(gs.Player.HP)
	maxhp := strconv.Itoa(gs.Player.Vit * (gs.Player.Res / 2))
	name := gs.Player.Name
	gold := strconv.Itoa(gs.Player.Gold)
	statusbar := []string{name, " | HP: ", hp, "/", maxhp, " | Gold: ", gold}
	DrawText(0, gs.ScreenHeight-1, strings.Join(statusbar, ""))
}

func (gs *GameState) DisplayMap() {
	for y := 0; y < gs.GameMap.Height; y++ {
		for x := 0; x < gs.GameMap.Width; x++ {
			tile := gs.GameMap.Data[x][y]
			//if tile.IsVisible == true {
			DrawRichText(x, y, tile.Char, tile.Color, termbox.ColorBlack)
			//}
		}
	}
}
