package main

type Character struct {
	X         int
	Y         int
	Char      string
	ClassID   int
	RaceID    int
	ClassName string
	RaceName  string
	Vit       int
	End       int
	Str       int
	Res       int
	Fai       int
	Wis       int
	Per       int
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
	{Key: "Pyromancer", Value: []int{18, 10, 16, 12, 8, 14, 12}},
	{Key: "Sorceror", Value: []int{16, 10, 14, 12, 6, 24, 8}},
	{Key: "Noble", Value: []int{38, 6, 6, 6, 6, 6, 6}}}

var RaceBase []Pair = []Pair{
	//   VIT|END|STR|RES|FAI|WIS|PER
	{Key: "Western Pilgrim", Value: []int{0, 0, 2, 0, 2, -2, -2}},
	{Key: "Northern Hallowman", Value: []int{2, 2, -2, 0, 0, 0, -2}},
	{Key: "Swamp Denizen", Value: []int{0, -2, 0, 0, -2, 2, 2}}}

func (ch Character) moveCharacter(x int, y int) {
	ch.X = x
	ch.Y = y
}

func (ch *Character) MoveLeft(gm *GameState) {
	if ch.X == 0 {
		return
	}
	if gm.GameMap.LocateTile(ch.X-1, ch.Y).IsWalkable {
		ch.X -= 1
	}
}

func (ch *Character) MoveRight(gm *GameState) {
	if ch.X == gm.GameMap.Width {
		return
	}
	if gm.GameMap.LocateTile(ch.X+1, ch.Y).IsWalkable {
		ch.X += 1
	}
}

func (ch *Character) MoveUp(gm *GameState) {
	if ch.Y == 0 {
		return
	}
	if gm.GameMap.LocateTile(ch.X, ch.Y-1).IsWalkable {
		ch.Y -= 1
	}
}

func (ch *Character) MoveDown(gm *GameState) {
	if ch.Y == gm.GameMap.Height {
		return
	}
	if gm.GameMap.LocateTile(ch.X, ch.Y+1).IsWalkable {
		ch.Y += 1
	}
}

func (character *Character) CreateChar(class int, race int) *Character {
	character.X = 2
	character.Y = 2

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

	return character
}
