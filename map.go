package main

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	"time"
)

const (
	MAXROOMW = 7 // Actually 15
	MINROOMW = 8
	MAXROOMH = 7 // Actually 15
	MINROOMH = 8
	MAXROOMS = 15 // Actually 25
	MINROOMS = 10

	WALL_CHAR       = "#"
	FLOOR_CHAR      = "."
	UPSTAIRS_CHAR   = ">"
	DOWNSTAIRS_CHAR = "<"
)

type Tile struct {
	Color       termbox.Attribute
	IsVisible   bool
	IsWalkable  bool
	BlocksLight bool
	IsWall      bool
	Char        string
	IsExit      bool
}

type Rect struct {
	StartX int
	StartY int
	Width  int
	Height int
}

type Map struct {
	Width  int
	Height int
	Data   [][]*Tile
	StartX int
	StartY int
}

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

func (gs *GameState) GenRooms() {
	rand.Seed(time.Now().Unix())
	width := rand.Intn(MAXROOMW) + MINROOMW
	height := rand.Intn(MAXROOMH) + MINROOMH

	level := gs.GameMap

	cells := []Rect{}

	for y := 0; y < gs.GameMap.Height; y += 10 {
		for x := 0; x < gs.GameMap.Width; x += 10 {
			cells = append(cells, Rect{StartX: x, StartY: y, Width: 10, Height: 10})
		}
	}

	for i, e := range cells {
		print(i)
		switch i {
		case 0:
			level.StartX = e.StartX + (width / 2)
			level.StartY = e.StartY + (height / 2)
			level.Data[level.StartX][level.StartY].Char = UPSTAIRS_CHAR
			level.Data[level.StartX][level.StartY].Color = termbox.ColorBlue
		}

		for y := e.StartY; y < height; y++ {
			for x := e.StartX; x < width; x++ {
				if x == width-1 || x == 0 { // Is it a wall?
					level.Data[x][y].IsWalkable = false
					level.Data[x][y].IsWall = true
				} else if y == height-1 || y == 0 { // Is it a wall?
					level.Data[x][y].IsWalkable = false
					level.Data[x][y].IsWall = true
				} else {
					level.Data[x][y].IsWalkable = true
					level.Data[x][y].IsWall = false
					level.Data[x][y].Char = FLOOR_CHAR
					level.Data[x][y].Color = termbox.ColorGreen
				}
			}
		}
	}
}
