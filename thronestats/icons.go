package thronestats

import (
	"fmt"
)


func GetCharacterIcon(character int, bskin bool) string {
	bskinTxt := ""

	if bskin {
		bskinTxt = "B"
	}

	return fmt.Sprintf("img/characters/sprMutant%d%sIdle.gif", character, bskinTxt)
}
