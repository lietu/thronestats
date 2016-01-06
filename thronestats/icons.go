package thronestats

import (
	"fmt"
)

var unsupportedEnemies = []int{
	20, 42, 43, 85, 87, 88,
}

var unsupportedWeapons = []int{
	63, 64, 72, 78, 79, 80, 85, 87, 88, 95, 106, 108, 120, 124, 201,
}

func GetEnemyIcon(enemy int) string {
	if IsIn(unsupportedEnemies, enemy) {
		return ""
	}

	return fmt.Sprintf("img/causesOfDeath/%d.gif", enemy)
}

func GetWeaponIcon(weapon int) string {
	if IsIn(unsupportedWeapons, weapon) {
		return ""
	}

	return fmt.Sprintf("img/weaponChoices/%d.gif", weapon)
}

func GetCrownIcon(crown int) string {
	return fmt.Sprintf("img/crownChoices/%d.gif", crown)
}

func GetMutationIcon(mutation int) string {
	return fmt.Sprintf("img/mutationChoices/%d.png", mutation)
}

func GetUltraIcon(character int, ultra int) string {
	ultraTxt := ""

	if ultra == 1 {
		ultraTxt = "A"
	} else if ultra == 2 {
		ultraTxt = "B"
	} else if ultra == 3 {
		ultraTxt = "C"
	} else {
		return ""
	}

	return fmt.Sprintf("img/ultras/%d%s.png", character, ultraTxt)
}

func GetCharacterIcon(character int, bskin bool) string {
	bskinTxt := ""

	if bskin {
		bskinTxt = "B"
	}

	return fmt.Sprintf("img/characters/sprMutant%d%sIdle.gif", character, bskinTxt)
}
