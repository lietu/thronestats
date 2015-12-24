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
	RunTime       int
	Kills         int
	IsDaily       bool
	IsWeekly      bool
	Ultra         bool
	LastDamagedBy int

	Area          int
	World         int
	Loop          int
	Level         string
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

func (rd *RunData) ReadFromApiResponse(ar *ApiResponse) {
	rd.Health = ToInt(ar.Health)
	rd.BSkin = (ar.BSkin == "1")
	rd.Character = ToInt(ar.Character)
	rd.Crown = ToInt(ar.Crown)
	rd.Mutations = parseMutations(ar.Mutations)
	rd.Weapons = parseWeapons(ToInt(ar.Weapon1), ToInt(ar.Weapon2))
	rd.RunTime = ToInt(ar.RunTime)
	rd.Kills = ToInt(ar.Kills)
	rd.IsDaily = (ar.IsDaily == "1")
	rd.IsWeekly = (ar.IsWeekly == "1")
	rd.Ultra = (ar.Ultra == "1")
	rd.LastDamagedBy = ToInt(ar.LastDamagedBy)

	rd.Area = ToInt(ar.Area)
	rd.World = ToInt(ar.World)
	rd.Loop = ToInt(ar.Loop)

	rd.Level = fmt.Sprintf("L%d %d-%d", rd.Loop, rd.World, rd.Area)
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
		false,
		false,
		false,
		0,
		0,
		0,
		0,
		"",
	}

	return &rd
}