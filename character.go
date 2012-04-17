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

var ClassBase map[string][]int = map[string][]int{
	//   VIT|END|STR|RES|FAI|WIS|PER
	"Knight":     []int{22, 10, 18, 14, 10, 8, 8},
	"Cleric":     []int{16, 10, 14, 12, 24, 6, 8},
	"Barbarian":  []int{24, 10, 24, 16, 6, 2, 8},
	"Pyromancer": []int{18, 10, 16, 12, 8, 14, 12},
	"Sorceror":   []int{16, 10, 14, 12, 6, 24, 8},
	"Noble":      []int{38, 6, 6, 6, 6, 6, 6},
	"Infested":   []int{13, 12, 13, 13, 13, 13, 13}}

var RaceBase map[string][]int = map[string][]int{
	//   VIT|END|STR|RES|FAI|WIS|PER
	"Western Pilgrim":    []int{0, 0, 2, 0, 2, -2, -2},
	"Northern Hallowman": []int{2, 2, -2, 0, 0, 0, -2},
	"Swamp Denizen":      []int{0, -2, 0, 0, -2, 2, 2}}

func (ch Character) moveCharacter(x int, y int) {
	ch.X = x
	ch.Y = y
}

func (ch Character) MoveLeft(gm *GameState) {
	if gm.GameMap.LocateTile(ch.X-1, ch.Y).IsWalkable {
		ch.X -= 1
	}
}

func (ch Character) MoveRight(gm *GameState) {
	if gm.GameMap.LocateTile(ch.X+1, ch.Y).IsWalkable {
		ch.X += 1
	}
}

func (ch Character) MoveUp(gm *GameState) {
	if gm.GameMap.LocateTile(ch.X, ch.Y-1).IsWalkable {
		ch.Y -= 1
	}
}

func (ch Character) MoveDown(gm *GameState) {
	if gm.GameMap.LocateTile(ch.X, ch.Y+1).IsWalkable {
		ch.Y += 1
	}
}

func (character Character) CreateChar(class int, race int) *Character {
	i := 0
	for name, info := range ClassBase {
		if i == class {
			character.ClassName = name
			character.ClassID = class
			character.Vit = info[0]
			character.End = info[1]
			character.Str = info[2]
			character.Res = info[3]
			character.Fai = info[4]
			character.Wis = info[5]
			character.Per = info[6]
		}
		i += 1
	}
	i = 0
	for name, info := range RaceBase {
		if i == race {
			character.RaceName = name
			character.RaceID = race
			character.Vit += info[0]
			character.End += info[1]
			character.Str += info[2]
			character.Res += info[3]
			character.Fai += info[4]
			character.Wis += info[5]
			character.Per += info[6]
		}
	}
	return &character
}
