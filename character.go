package main

type Character struct {
	X         int
	Y         int
	Char      string
	Name      string
	ClassID   int
	RaceID    int
	ClassName string
	RaceName  string
	HP        int
	Vit       int
	End       int
	Str       int
	Res       int
	Fai       int
	Wis       int
	Per       int
	Gold      int
	Inventory []*Item
	HEAD      *Item
	CHEST     *Item
	ARMS      *Item
	HANDS     *Item
	LEGS      *Item
	FEET      *Item
}

type Pair struct {
	Key   string
	Value []int
}

var ClassBase []Pair = []Pair{
	//   VIT|END|STR|RES|FAI|WIS|PER
	{Key: "Knight", Value: []int{22, 10, 18, 14, 10, 8, 8}},
	{Key: "Cleric", Value: []int{16, 10, 14, 12, 24, 6, 8}},
	{Key: "Barbarian", Value: []int{24, 10, 24, 16, 6, 2, 8}},
	{Key: "Alchemist", Value: []int{18, 10, 16, 12, 8, 14, 12}},
	{Key: "Sorceror", Value: []int{16, 10, 14, 12, 6, 24, 8}},
	{Key: "Nobility", Value: []int{38, 6, 6, 6, 6, 6, 6}}}

var RaceBase []Pair = []Pair{
	//                                  VIT|END|STR|RES|FAI|WIS|PER
	{Key: "Orc",            Value: []int{0, 0, 2, 0, 2, -2, 0}},
	{Key: "Dwarf",          Value: []int{2, 2, -2, 0, 0, 0, -2}},
	{Key: "Octavian",       Value: []int{0, -2, 0, 0, -2, 2, 2}}}

func (ch Character) moveCharacter(x int, y int) {
	ch.X = x
	ch.Y = y
}

func (ch *Character) MoveLeft(gm *GameState) {
	if ch.X == 0 {
		return
	}
	tile, exists := gm.GameMap.LocateTile(ch.X-1, ch.Y)
	if exists == true && tile.IsWalkable {
		ch.X -= 1
	}
}

func (ch *Character) MoveRight(gm *GameState) {
	if ch.X == gm.GameMap.Width {
		return
	}
	tile, exists := gm.GameMap.LocateTile(ch.X+1, ch.Y)
	if exists == true && tile.IsWalkable {
		ch.X += 1
	}
}

func (ch *Character) MoveUp(gm *GameState) {
	if ch.Y == 0 {
		return
	}
	tile, exists := gm.GameMap.LocateTile(ch.X, ch.Y-1)
	if exists == true && tile.IsWalkable {
		ch.Y -= 1
	}
}

func (ch *Character) MoveDown(gm *GameState) {
	if ch.Y == gm.GameMap.Height {
		return
	}
	tile, exists := gm.GameMap.LocateTile(ch.X, ch.Y+1)
	if exists == true && tile.IsWalkable {
		ch.Y += 1
	}
}

func (character *Character) CreateChar(class int, race int, name string, gs *GameState) *Character {
	character.X = gs.GameMap.StartX
	character.Y = gs.GameMap.StartY

	character.Name = name

	classInfo := ClassBase[class]
	character.ClassName = classInfo.Key
	character.ClassID = class
	character.Vit = classInfo.Value[0]
	character.End = classInfo.Value[1]
	character.Str = classInfo.Value[2]
	character.Res = classInfo.Value[3]
	character.Fai = classInfo.Value[4]
	character.Wis = classInfo.Value[5]
	character.Per = classInfo.Value[6]

	raceInfo := RaceBase[race]
	character.RaceName = raceInfo.Key
	character.RaceID = race
	character.Vit += raceInfo.Value[0]
	character.End += raceInfo.Value[1]
	character.Str += raceInfo.Value[2]
	character.Res += raceInfo.Value[3]
	character.Fai += raceInfo.Value[4]
	character.Wis += raceInfo.Value[5]
	character.Per += raceInfo.Value[6]

	character.HP = character.Vit * (character.Res / 2)
	character.Gold = 5

	return character
}
