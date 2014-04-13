package main

import (
	//. "github.com/chownplusx/Termbox-Additions"
	"github.com/Chownie/Termbox-Additions/utils"
	"github.com/nsf/termbox-go"
	"strconv"
	"strings"

//	"fmt"
)

const (
    SPLASH = "  ▄▀  ████▄ █▄▄▄▄   ▄▀  ████▄    ▄  \n▄▀    █   █ █  ▄▀ ▄▀    █   █     █ \n█ ▀▄  █   █ █▀▀▌  █ ▀▄  █   █ ██   █\n█   █ ▀████ █  █  █   █ ▀████ █ █  █\n ███          █    ███        █  █ █\n             ▀                █   ██"
    // Offsets
    offset = 1
)

func (gs *GameState) DisplaySplash(fg termbox.Attribute) {
	x := (gs.ScreenWidth / 2)
    offsetx := len(strings.Split(SPLASH, "\n")[0])/4
    offsety := len(strings.Split(SPLASH, "\n"))/4
	y := (gs.ScreenHeight / 2)
	utils.DrawRichTextMulti(x-offsetx, y-offsety, SPLASH, fg, termbox.ColorDefault)
	utils.DrawText(x-(len("Press any key to start")/2), y+(y/2), "Press any key to start")
}

func (ch *Character) DisplayPlayer() {
	utils.DrawRichText(ch.X+offset, ch.Y+offset, "@", termbox.ColorGreen, termbox.ColorBlack)
}

func (gs *GameState) DisplayStatus() {
	HP := strconv.Itoa(gs.Player.HP)
	MaxHP := strconv.Itoa(gs.Player.Vit * (gs.Player.Res / 2))
    END := strconv.Itoa(gs.Player.End)
    STR := strconv.Itoa(gs.Player.Str)
    RES := strconv.Itoa(gs.Player.Res)
    FAI := strconv.Itoa(gs.Player.Fai)
    WIS := strconv.Itoa(gs.Player.Wis)
    PER := strconv.Itoa(gs.Player.Per)
	gold := strconv.Itoa(gs.Player.Gold)
    name_hp := []string{    "[" + gs.Player.Name + " the " + gs.Player.RaceName + " " + gs.Player.ClassName + "]",
                            "HP: " + HP + "/" + MaxHP + " Gold: " + gold + "¤"  }

    stats := []string { "End: " + END + " Str: " + STR + " Res: " + RES,
                        "Fai: " + FAI + " Wis: " + WIS + " Per: " + PER }
    y := gs.GameMap.Height+offset+5
    width := gs.GameMap.Width-(offset*2)
    length := 5
    utils.DrawBox(0, y, width, length, utils.CONNECT_TOP)
    utils.DrawTextMulti(offset+1, y+2, strings.Join(name_hp, "\n"))
    utils.DrawTextMulti(width/2 + width/5, y+2, strings.Join(stats, "\n"))
}

func (gs *GameState) DisplayMessages() {
    y := gs.GameMap.Height+offset
    width := gs.GameMap.Width-(offset*2)
    length := 5
    utils.DrawBox(0, y, width, length, utils.CONNECT_TOP)
    utils.DrawTextMulti(offset+1, y+1, strings.Join(*gs.Messages,"\n"))
}

func (gs *GameState) DisplayInventory() {
    utils.DrawBox(gs.GameMap.Width+offset, 0, 20, gs.GameMap.Height+offset, utils.CONNECT_LEFT+utils.CONNECT_BOT)
    teststring := []string{"Inventory here", "soon."}
    utils.DrawTextMulti(gs.GameMap.Width+offset+2, 1, strings.Join(teststring, "\n"))
}

func (gs *GameState) DisplayMap() {
    utils.DrawBox(0,0,gs.GameMap.Width-(offset*2), gs.GameMap.Height+offset, utils.CONNECT_BOT)
    for y := 0; y < gs.GameMap.Height; y++ {
		for x := 0; x < gs.GameMap.Width; x++ {
			tile := gs.GameMap.Data[x][y]
			//if tile.IsVisible == true {
			utils.DrawRichText(x+offset, y+offset, tile.Char, tile.Color, termbox.ColorBlack)
			//}
		}
	}
}
