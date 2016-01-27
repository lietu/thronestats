package thronestats

import (
	"fmt"
	"unicode/utf8"
)


type RunData struct {
	Health        int
	BSkin         bool
	Character     int
	Crown         int
	Mutations     []int
	Weapons       []int
	Timestamp     int
	Kills         int
	Ultra         int
	LastDamagedBy int

	Area          int
	World         int
	Loop          int
	Level         string
}

func NewRunData() *RunData {
	rd := RunData{
		0,
		false,
		0,
		0,
		[]int{},
		[]int{},
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		"",
	}

	return &rd
}

type RunDataContainer struct {
	Current *RunData
	Previous *RunData
}

func NewRunDataContainer() *RunDataContainer {
	rdc := RunDataContainer{
		NewRunData(),
		NewRunData(),
	}

	return &rdc
}

func parseMutations(mutationText string) []int {
	mutations := []int{}

	one, _ := utf8.DecodeRuneInString("1")

	for key, value := range mutationText {
		if value == one {
			mutations = append(mutations, key)
		}
	}

	return mutations
}

func parseWeapons(weapon1 int, weapon2 int) []int {
	if weapon1 > 0 {
		if weapon2 > 0 {
			if weapon1 < weapon2 {
				return []int{weapon1, weapon2}
			} else {
				return []int{weapon2, weapon1}
			}
		} else {
			return []int{weapon1}
		}
	} else if weapon2 > 0 {
		return []int{weapon2}
	}

	return []int{}
}

func (rd *RunData) ReadFromApiResponseRun(arr *ApiResponseRun) {
	rd.Health = arr.Health
	rd.BSkin = (arr.BSkin == 1)
	rd.Character = arr.Character
	rd.Crown = arr.Crown
	rd.Mutations = parseMutations(arr.Mutations)
	rd.Weapons = parseWeapons(arr.Weapon1, arr.Weapon2)
	rd.Timestamp = arr.Timestamp
	rd.Kills = arr.Kills
	rd.Ultra = arr.Ultra
	rd.LastDamagedBy = arr.LastDamagedBy

	rd.Area = arr.Area
	rd.World = arr.World
	rd.Loop = arr.Loop

	rd.Level = fmt.Sprintf("L%d %d-%d", rd.Loop, rd.World, rd.Area)
}

func (rdc *RunDataContainer) ReadFromApiResponse(ar *ApiResponse) {
	if ar.Current != nil {
		rdc.Current.ReadFromApiResponseRun(ar.Current)
	}
	if ar.Previous != nil {
		rdc.Previous.ReadFromApiResponseRun(ar.Previous)
	}
}

func (rd *RunData) Print() {
	fmt.Printf("Character: %d\n", rd.Character)
	fmt.Printf("LastDamagedBy: %d\n", rd.LastDamagedBy)
	fmt.Printf("Level: %s\n", rd.Level)
	fmt.Printf("Crown: %d\n", rd.Crown)
	weapons := ""
	for _, id := range rd.Weapons {
		weapons = fmt.Sprintf("%s%d, ", weapons, id)
	}
	fmt.Printf("Weapons: %s\n", weapons)
	fmt.Printf("BSkin: %t\n", rd.BSkin)
	fmt.Printf("Ultra: %d\n", rd.Ultra)
	mutations := ""
	for _, id := range rd.Mutations {
		mutations = fmt.Sprintf("%s%d, ", mutations, id)
	}
	fmt.Printf("Mutations: %s\n", mutations)
	fmt.Printf("Kills: %d\n", rd.Kills)
	fmt.Printf("Health: %d\n", rd.Health)
	fmt.Printf("Timestamp: %d\n", rd.Timestamp)
}