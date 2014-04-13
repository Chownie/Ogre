package main

import (
	//"github.com/chownplusx/Termbox-Additions/utils"
	"github.com/nsf/termbox-go"
	//"math"
	//	"fmt"
	//"math/rand"
	"strconv"
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
	SCALEMAX = 6
	MINROOMS = 5
	MAXROOMS = 8
	// Defining directions
	BE_HORI = iota
	BE_VERT
	// Directions!
	DI_FOREWARD = -1
	DI_BACKWARD = 1
    //Upper and lower cutting bounds
    LOWER = 30
    UPPER = 50
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
    wl := new(Tile)
    wl.Char = WALL_CHAR
    wl.IsWalkable = false
    wl.IsWall = true
    wl.IsVisible = false
    wl.Color = termbox.ColorCyan
    gs.FillArea(X,   Y,     1, H, wl)   // fill Left wall
    gs.FillArea(X+W, Y,     1, H+1, wl)   // Fill right wall
    gs.FillArea(X,   Y,     W, 1, wl)   // Fill top wall
    gs.FillArea(X,   Y+H,   W, 1, wl)   // Fill bottom wall
	return &room
}

func (gs *GameState) AddRoomToGrid(rm *Room) {
	gs.GameMap.RoomList = append(gs.GameMap.RoomList, rm)
}

// Randomly generates a number between Min and Max-1
// So Random(2, 5) can generate 2, 3 and 4 but NOT 5.
func (gs *GameState) Random(min, max int) int {
	if min > max {
		panic("Min: " + strconv.Itoa(min) + " Max: " + strconv.Itoa(max))
	}
	return gs.RNG.Intn((max+1)-min) + min
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

//Fills a specific room with one particular tile type.
func (gs *GameState) FillRoom(rm *Room, tl *Tile) {
	gs.FillArea(rm.StartX, rm.StartY, rm.Width, rm.Height, tl)
}

// Fills a rect with the given tile type
func (gs *GameState) FillArea(X, Y, W, H int, tl *Tile) {
	for y := Y; y < (Y + H); y++ {
		for x := X; x < (X + W); x++ {
			tile, exists := gs.GameMap.LocateTile(x, y)
			if !exists {
				panic("out of bounds!\nX: " + strconv.Itoa(x) + " Y: " + strconv.Itoa(y))
			}
			tile.Char = tl.Char
            tile.IsWalkable = tl.IsWalkable
            tile.IsWall = tl.IsWall
            tile.Color = tl.Color

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
				return false
			}
		}
		return true
	}
	return false
}

// Controls the generation of the rooms itself.
// Using a Binary Space Partition system to gen the map
func (gs *GameState) GenMap() {
	// I've attempted this many times, always with messy code and wierd logic,
	// at 2AM or on a train while sleepless. Hopefully this attempt will go better
	// with some documentation included.
	// Attempt No: 7 /6 /5 /4 /3 /2

	gs.BSPartition(0, 0, gs.GameMap.Width-1, gs.GameMap.Height-1, 3)
    gs.Player.X = 5
    gs.Player.Y = 5
}

// Recursively creates the BSP
func (gs *GameState) BSPartition(X, Y, W, H, Count int) {
    //Bound := gs.Random(LOWER, UPPER)
	Bound := 50
    if Count > 0 {
		direction := gs.Random(0, 1)
		switch direction {
		case 0: // Cut the map into 2 vertically
			//posx := gs.Random(X, X+W)
            posx := ((X+W)*Bound)/100

            gs.BSPartition(X,       Y,      posx,       H,      Count-1)   //Left
			gs.BSPartition(posx,    Y,      W-posx,     H,      Count-1)   //Right
		case 1: // Cut the map into 2 Horizontally
			//posy := gs.Random(Y, Y+H)
            posy := ((Y+H)*Bound)/100


            gs.BSPartition(X,       Y,      W,          posy,    Count-1)    //Top
			gs.BSPartition(X,       posy,   W,          H-posy,  Count-1)    //Bottom
		default:
			panic("Direction: " + strconv.Itoa(direction))
		}
	} else {
		RoomWidth := gs.Random(SCALEMIN, W)
		RoomHeight := gs.Random(SCALEMIN, H)
		
        rm := gs.NewRoom(X, Y, RoomWidth, RoomHeight, false, false)
		gs.AddRoomToGrid(rm)
		return
	}
	return
}

// Generates Rooms only
func (gs *GameState) MakeRoom(X, Y, W, H, direction, axis int) bool {
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
