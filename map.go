package main

import (
	//"github.com/chownplusx/Termbox-Additions/utils"
	"github.com/nsf/termbox-go"
	//"math"
	//"fmt"
	"math/rand"
	"time"
	//	"strconv"
)

const (
	// Characters denoting objects (duh)
	WALL_CHAR       = "#"
	FLOOR_CHAR      = "."
	UPSTAIRS_CHAR   = ">"
	DOWNSTAIRS_CHAR = "<"
)

// Contains a "stack" of game maps to cycle
// through as different depths are reached.
type Level struct {
	Stack []*Map
}

// Contains the startX, startY (the stairs) positions,
// plus the tile array and the map info such as width
// and height.
type Map struct {
	Width  int
	Height int
	Data   [][]*Tile
	StartX int
	StartY int
	Grid   []Section
}

type Tile struct {
	Color       termbox.Attribute
	IsVisible   bool
	IsWalkable  bool
	BlocksLight bool
	IsWall      bool
	Char        string
	IsExit      bool
}

type Section struct {
	StartX      int
	StartY      int
	Width       int
	Height      int
	Connected   bool
	ConnectedTo []Section
}

func Random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// Helper function, looks within the multi-dimensional slice
// and returns a tile object and a boolean. If the tile exists
// it will return said tile's object with "true". Should the tile
// not exist, it returns a blank tile object with "false"
func (level *Map) LocateTile(x, y int) (*Tile, bool) {
	retTile := &Tile{}
	exists := false
	if x > level.Width-1 || x < 0 {
		retTile = &Tile{}
		exists = false
	} else if y > level.Height-1 || y < 0 {
		retTile = &Tile{}
		exists = false
	} else {
		retTile = level.Data[x][y]
		exists = true
	}
	return retTile, exists
}

// Fills the map with "blank" tile objects. Here means
// non-visible, non-walkable walls.
func BlankMap(width, height int) *Map {
	final := [][]*Tile{}
	temp := []*Tile{}

	for y := 0; y < height; y++ {
		temp = []*Tile{}
		for x := 0; x < width; x++ {
			tile := Tile{}
			tile.Color = termbox.ColorCyan
			tile.IsWalkable = false
			tile.IsVisible = false
			tile.IsWall = true
			tile.IsExit = false
			tile.Char = WALL_CHAR
			temp = append(temp, &tile)
		}
		final = append(final, temp)
	}
	return &Map{Width: width, Height: height, Data: final}
}

//Fills a section with one particular tile type.
func (gs *GameState) FillSection(sc *Section, tl *Tile) {
	// termbox.Close()

	// fmt.Println("BEING CALLED")
	for y := sc.StartY; y < (sc.StartY + sc.Height); y++ {
		for x := sc.StartX; x < (sc.StartX + sc.Width); x++ {
			// fmt.Println("X: ", x)
			// fmt.Println("Y: ", y)
			tile, exists := gs.GameMap.LocateTile(x, y)
			if !exists {
				
			}
			tile.SetFloor()
		}
	}
}

// Helper function. Sets a tile to the floortype.
func (tl *Tile) SetFloor() {
	tl.BlocksLight = false
	tl.Color = termbox.ColorGreen
	tl.IsWalkable = true
	tl.IsWall = false
	tl.Char = FLOOR_CHAR
}

// Controls the generation of the rooms itself.
// Cuts the map into a small grid.
func (gs *GameState) GenRooms() {
	// I've attempted this 4 or 5 times, always with messy code and wierd logic,
	// at 2AM or on a train while sleepless. Hopefully this attempt will go better
	// with some documentation included.

	// Cut the map into squares
	// I do this by dividing the width/height of the map by a magic number
	// This gives us the number of squares
	scale := 6
	gridsize := (gs.GameMap.Height) / scale

	for i := 0; i < (scale*scale); i++ {
		StartX := 0
		StartY := 0

		StartY = (gs.GameMap.Height / scale) * (i / scale)
		if i%scale == 0 {
			StartX = (gs.GameMap.Width / scale) * (scale - 1)
		} else {
			StartX = (gs.GameMap.Width / scale) * (i % scale) - gridsize
		}

		height := 6
		width := 6

		StartX += 2
		StartY += 2

		var Connected bool
		if i == 1 {
			gs.Player.X = StartX + (width/2)
			gs.Player.Y = StartY + (height/2)
			Connected = true
		} else {
			Connected = false
		}

		Room := new(Section)
		// fmt.Println("NEW SECTION")
		Room.Connected = Connected
		Room.StartX = StartX
		Room.StartY = StartY
		Room.Height = height
		Room.Width = width
		gs.GameMap.Grid = append(gs.GameMap.Grid, *Room)

		gs.GenRoute()
	}
}

func (gs *GameState) GenRoute() {
	//Split into two loops because they do different things
	for i := 0; i < len(gs.GameMap.Grid); i++ {
		tl := new(Tile)
		tl.SetFloor()
		gs.FillSection(&gs.GameMap.Grid[i], tl)
	}
}
