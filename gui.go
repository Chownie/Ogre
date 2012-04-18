package main

import (
	//. "github.com/chownplusx/Termbox-Additions"
	. "github.com/chownplusx/Termbox-Additions/utils"
	"github.com/nsf/termbox-go"
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

func (gs *GameState) DisplayMap() {
	for x := 0; x < gs.GameMap.Width; x++ {
		for y := 0; y < gs.GameMap.Height; y++ {
			tile := gs.GameMap.Data[x][y]
			if tile.IsVisible == true {
				DrawRichText(x, y, tile.Char, tile.Color, termbox.ColorBlack)
			}
		}
	}
}
