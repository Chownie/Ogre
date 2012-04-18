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
)

type Tile struct {
	Color       termbox.Attribute
	IsVisible   bool
	IsWalkable  bool
	BlocksLight bool
	Char        string
	IsExit      bool
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
			tile.IsWalkable = true
			tile.IsVisible = false
			tile.IsExit = false
			tile.Char = "."
			temp = append(temp, &tile)
		}
		final = append(final, temp)
	}
	return &Map{Width: width, Height: height, Data: final}
}

func (level *Map) GenerateMap() {
	rand.Seed(time.Now().Unix())
	roomCount := rand.Intn(MAXROOMS) + MINROOMS
	print(roomCount)
}

func (level *Map) MakeRoom(roomcount int, exit bool) {
	rand.Seed(time.Now().Unix())
	width := rand.Intn(MAXROOMW) + MINROOMW
	height := rand.Intn(MAXROOMH) + MINROOMH

	switch roomcount {
	case 0:
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if x == width-1 || x == 0 {
					level.Data[x+10][y].Char = "#"
					level.Data[x+10][y].IsWalkable = false
				} else if y == height-1 || y == 0 {
					level.Data[x+10][y].Char = "#"
					level.Data[x+10][y].IsWalkable = false
				}
				level.Data[x+10][y].Color = termbox.ColorGreen
				level.Data[x+10][y].IsExit = false
			}
		}
	}
}
