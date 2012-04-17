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
}

func (level *Map) LocateTile(x, y int) *Tile {
	return level.Data[x][y]
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
			tile.IsVisible = true
			tile.IsExit = false
			tile.Char = "."
			temp = append(temp, &tile)
		}
		final = append(final, temp)
	}
	return &Map{Width: width, Height: height, Data: final}
}

func (level Map) MakeRoom(roomcount int, exit bool) bool {
	rand.Seed(time.Now().Unix())
	width := rand.Intn(MAXROOMW) + MINROOMW
	height := rand.Intn(MAXROOMH) + MINROOMH

	switch roomcount {
	case 0:
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				level.Data[x][y].Color = termbox.ColorCyan
				level.Data[x][y].IsWalkable = true
				level.Data[x][y].IsVisible = true
				level.Data[x][y].IsExit = false
			}
		}
	}
	return true
}
