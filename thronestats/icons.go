package thronestats

import (
	"fmt"
)

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
