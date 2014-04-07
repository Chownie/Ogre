package main

import (
	//"github.com/chownplusx/Termbox-Additions/utils"
	"github.com/nsf/termbox-go"
	//"math"
	//	"fmt"
	"math/rand"
	//"strconv"
)

const (
	// Characters denoting objects
	UNSET_CHAR      = ""
	WALL_CHAR       = "#"
	FLOOR_CHAR      = "."
	UPSTAIRS_CHAR   = ">"
	DOWNSTAIRS_CHAR = "<"
	// Let's define some stuff about the map
	SCALEMIN = 3
	SCALEMAX = 7
	MINROOMS = 5
	MAXROOMS = 8
	// Defining directions
	BE_HORI = iota
	BE_VERT
	// Directions!
	DI_FOREWARD = -1
	DI_BACKWARD = 1
)

// Contains a "stack" of game maps to cycle
// through as different depths are reached.
type Level struct {
	Stack []*Map
}

type Tile struct {
	Color       termbox.Attribute
	IsVisible   bool
	IsWalkable  bool
	BlocksLight bool
	IsWall      bool
	Char        string
}

// Contains the startX, startY (the stairs) positions,
// plus the tile array and the map info such as width
// and height.
type Map struct {
	Width    int
	Height   int
	Data     [][]*Tile
	StartX   int
	StartY   int
	RoomList []*Room
}

type Room struct {
	StartX    int
	StartY    int
	Width     int
	Height    int
	Connected bool
	Special   bool
}

func (gs *GameState) NewRoom(X, Y, W, H int, Connect, Spec bool) *Room {
	room := Room{StartX: X,
		StartY:    Y,
		Width:     W,
		Height:    H,
		Connected: Connect,
		Special:   Spec}
	tl := new(Tile)
	tl.SetFloor()
	gs.FillRoom(&room, tl)
	return &room
}

func (gs *GameState) AddRoomToGrid(rm *Room) {
	gs.GameMap.RoomList = append(gs.GameMap.RoomList, rm)
}

func (gs *GameState) Random(min, max int) int {
	return rand.Intn(max-min) + min
}

// Helper function, looks within the multi-dimensional slice
// and returns a tile object and a boolean. If the tile exists
// it will return said tile's object with "true". Should the tile
// not exist, it returns a blank tile object with "false"
func (level *Map) LocateTile(x, y int) (*Tile, bool) {
	// fmt.Println(x, y)
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

	for x := 0; x < width; x++ {
		temp = []*Tile{}
		for y := 0; y < height; y++ {
			tile := Tile{}
			tile.Color = termbox.ColorCyan
			tile.IsWalkable = false
			tile.IsVisible = false
			tile.IsWall = true
			tile.Char = UNSET_CHAR
			temp = append(temp, &tile)
		}
		final = append(final, temp)
	}
	return &Map{Width: width, Height: height, Data: final}
}

//Fills a section with one particular tile type.
func (gs *GameState) FillRoom(rm *Room, tl *Tile) {
	for y := rm.StartY; y < (rm.StartY + rm.Height); y++ {
		for x := rm.StartX; x < (rm.StartX + rm.Width); x++ {
			// fmt.Println("X: ", x)
			// fmt.Println("Y: ", y)

			tile, exists := gs.GameMap.LocateTile(x, y)
			if !exists {
				panic("out of bounds")
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

// Verifies that the selected area is "clear" to map on
func (gs *GameState) VerifyArea(X, Y, W, H int) bool {
	for y := Y; y < Y+H; y++ {
		for x := X; x < X+W; x++ {
			tl, exists := gs.GameMap.LocateTile(x, y)
			if exists == false {
				return false
			} else if tl.Char != UNSET_CHAR {
				return false
			}
		}
	}
	return true
}

// Verifies that a single line is "clear" to map on
func (gs *GameState) VerifyLine(X, Y, L, Direction, Axis int) bool {
	if Axis == BE_HORI {
		for x := X; x != X+(L*Direction); x += Direction {
			tl, exists := gs.GameMap.LocateTile(x, Y)
			if exists == false {
				return false
			} else if tl.Char != UNSET_CHAR {
				Loggy.Println("Tile is ", tl.Char)
				return false
			}
		}
		return true
	} else if Axis == BE_VERT {
		for y := Y; y != Y+(L*Direction); y += Direction {
			tl, exists := gs.GameMap.LocateTile(X, y)
			if exists == false {
				return false
			} else if tl.Char != UNSET_CHAR {
				Loggy.Println("CurrRoom X: ", X, " CurrRoom Y: ", Y)
				Loggy.Println(gs.GameMap.RoomList[len(gs.GameMap.RoomList)-1])
				return false
			}
		}
		return true
	}
	return false
}

// Controls the generation of the rooms itself.
// Starts with 1 room in the center, picks a random wall
// and checks if there's enough space on the other side for a wall/corridor
// if there is, cut a tile out and replace it with a door
// then draw a corridor and make another room :)
func (gs *GameState) GenMap() {
	// I've attempted this many times, always with messy code and wierd logic,
	// at 2AM or on a train while sleepless. Hopefully this attempt will go better
	// with some documentation included.
	// Attempt No: 6! /5 /4 /3 /2

	originW := gs.Random(SCALEMIN, SCALEMAX)
	originH := gs.Random(SCALEMIN, SCALEMAX)
	rm := gs.NewRoom((gs.GameMap.Width/2)-(originW/2),
		(gs.GameMap.Height/2)-(originH/2),
		originW,
		originH,
		false,
		false)
	gs.AddRoomToGrid(rm)

	gs.GameMap.StartX = rm.StartX + (rm.Width / 2)
	gs.GameMap.StartY = rm.StartY + (rm.Height / 2)

	gs.Player.X = gs.GameMap.StartX
	gs.Player.Y = gs.GameMap.StartY

	gs.GenRooms()
}

func (gs *GameState) GenRooms() {
	RoomCount := gs.Random(MINROOMS, MAXROOMS)
	for i := 0; i < RoomCount; i++ {
		Direction := gs.Random(1, 4)
		//newRoomW := gs.Random(SCALEMIN, SCALEMAX)
		//newRoomH := gs.Random(SCALEMIN, SCALEMAX)

		prevRoom := gs.GameMap.RoomList[len(gs.GameMap.RoomList)-1]

		open := true
		switch Direction {
		case 1: // UPWARD, Y goes down  (BE_VERT, DI_BACKWARD)
			point := gs.Random(prevRoom.StartX, prevRoom.StartX+prevRoom.Width)
			if i%2 == 0 {
				open = gs.MakeHall(point-1, prevRoom.StartY,
					DI_BACKWARD,
					BE_VERT)
			}
		case 2: // LEFT, X goes down (BE_HORI, DI_BACKWARD)
			point := gs.Random(prevRoom.StartY, prevRoom.StartY+prevRoom.Height)
			if i%2 == 0 {
				open = gs.MakeHall(prevRoom.StartX, point,
					DI_BACKWARD,
					BE_HORI)
			}
		case 3: // RIGHT, X goes up (BE_HORI, DI_FOREWARD)
			point := gs.Random(prevRoom.StartY, prevRoom.StartY+prevRoom.Height)
			if i%2 == 0 {
				open = gs.MakeHall(prevRoom.StartX+prevRoom.Width, point,
					DI_FOREWARD,
					BE_HORI)
			}
		case 4: // DOWNWARD, Y goes up (BE_VERT, DI_FOREWARD)
			point := gs.Random(prevRoom.StartX, prevRoom.StartX+prevRoom.Width)
			if i%2 == 0 {
				open = gs.MakeHall(point+1, prevRoom.StartY+prevRoom.Height,
					DI_FOREWARD,
					BE_HORI)
			}
		}
		if open == false {
			i -= 1
		}
	}
}

// Generates Rooms only
func (gs *GameState) MakeRoom(X, Y, W, H, direction, axis int) bool {

	//rm := gs.NewRoom(X,Y,W,H,false,false)
	//gs.AddRoomToGrid(rm)

	if axis == BE_VERT {
		//open := gs.VerifyArea(X, Y, W, H)
		//if open == false {
		//	return false
		//}
		rm := gs.NewRoom(X, Y, W, H, false, false)
		gs.AddRoomToGrid(rm)
		return true
	} else if axis == BE_HORI {
		open := gs.VerifyArea(X, Y, W, H)
		if open == false {
			return false
		}
		rm := gs.NewRoom(X, Y, W, H, false, false)
		gs.AddRoomToGrid(rm)
		return true
	}
	return false
}

// Generates the corridors between rooms
func (gs *GameState) MakeHall(startX, startY, direction, axis int) bool {
	L := gs.Random(SCALEMIN, SCALEMAX)

	if axis == BE_VERT {
		//open := gs.VerifyLine(startX, startY, L, direction, axis)
		//if open == false {
		//	return false
		//}
		rm := gs.NewRoom(startX, startY+(L*direction), 1, L, false, false)
		gs.AddRoomToGrid(rm)
		return true
	} else if axis == BE_HORI {
		open := gs.VerifyLine(startX, startY, L, direction, axis)
		if open == false {
			return false
		}
		rm := gs.NewRoom(startX+(L*direction), startY, L, 1, false, false)
		gs.AddRoomToGrid(rm)
		return true
	}
	return false
}
