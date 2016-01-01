package thronestats

import (
	"os"
	"fmt"
	"log"
	"strconv"
	"encoding/json"
	"io/ioutil"
)

var GlobalStats *StatsContainer

type StatsContainer struct {
	SteamId64              string           `json:"-"`
	Running                bool             `json:"-"`
	RunStats               RunStats         `json:"-"`
	Runs                   int              `json:"runs"`

	Characters             map[string]int   `json:"characters"`
	CausesOfDeath          map[string]int   `json:"causesOfDeath"`
	MutationChoices        map[string]int   `json:"mutationChoices"`
	WeaponChoices          map[string]int   `json:"weaponChoices"`
	CrownChoices           map[string]int   `json:"crownChoices"`

	DeathsByLevel          map[string]int   `json:"deathsByLevel"`

	MostPopularWeapon      string           `json:"mostPopularWeapon"`
	MostCommonCauseOfDeath string           `json:"mostCommonCauseOfDeath"`
	MostPopularMutation    string           `json:"mostPopularMutation"`
	MostPopularCrown       string           `json:"mostPopularCrown"`
	MostPopularCharacter   string           `json:"mostPopularCharacter"`
	MostCommonDeathLevel   string           `json:"mostCommonDeathLevel"`
}

func (sc *StatsContainer) getFilename() string {
	return fmt.Sprintf("data_%s.json", sc.SteamId64)
}

func (sc *StatsContainer) getMostPopular(items map[string]int, ignore string) string {
	lastCount := 0
	mostPopular := "0"

	for key, count := range items {
		if count > lastCount && key != ignore {
			lastCount = count
			mostPopular = key
		}
	}

	return mostPopular
}

func (sc *StatsContainer) getRate(runs int, currentRun bool) string {
	total := sc.Runs
	if currentRun {
		total += 1
	}
	return fmt.Sprintf("%.2f%%", (float64(runs) / float64(total)) * 100)
}

func (sc *StatsContainer) GetCrownRate(mutation int, extra int) string {
	return sc.getRate(sc.CrownChoices[strconv.Itoa(mutation)] + extra, extra == 1)
}

func (sc *StatsContainer) GetMutationRate(mutation int, extra int) string {
	return sc.getRate(sc.MutationChoices[strconv.Itoa(mutation)] + extra, extra == 1)
}

func (sc *StatsContainer) GetCharacterRate(character int, extra int) string {
	return sc.getRate(sc.WeaponChoices[strconv.Itoa(character)] + extra, extra == 1)
}

func (sc *StatsContainer) GetWeaponRate(weapon int, extra int) string {
	return sc.getRate(sc.WeaponChoices[strconv.Itoa(weapon)] + extra, extra == 1)
}

func (sc *StatsContainer) GetCauseOfDeathRate(enemy int) string {
	return sc.getRate(sc.CausesOfDeath[strconv.Itoa(enemy)], false)
}

func (sc *StatsContainer) GetLevelDeathRate(level string) string {
	return sc.getRate(sc.DeathsByLevel[level], false)
}

func (sc *StatsContainer) UpdateStats() {
	// Calculate stats
	weapon := sc.getMostPopular(sc.WeaponChoices, "1")
	causeOfDeath := sc.getMostPopular(sc.CausesOfDeath, "")
	mutation := sc.getMostPopular(sc.MutationChoices, "")
	crown := sc.getMostPopular(sc.CrownChoices, "1")
	character := sc.getMostPopular(sc.Characters, "0")
	level := sc.getMostPopular(sc.DeathsByLevel, "")

	sc.MostPopularWeapon = Weapons[ToInt(weapon)]
	sc.MostCommonCauseOfDeath = Enemies[ToInt(causeOfDeath)]
	sc.MostPopularMutation = Mutations[ToInt(mutation)]
	sc.MostPopularCrown = Crowns[ToInt(crown)]
	sc.MostPopularCharacter = Characters[ToInt(character)]
	sc.MostCommonDeathLevel = level
}

func (sc *StatsContainer) load() {
	filename := sc.getFilename()
	data, err := ioutil.ReadFile(filename)

	if err == nil {
		log.Printf("Loaded stats information from %s", filename)
		json.Unmarshal(data, sc)
	} else {
		if os.IsNotExist(err) {
			log.Printf("%s does not exist, do not have existing data.", filename)
			sc.Save()
		} else {
			log.Fatalf("error: %v", err)
		}
	}

	sc.UpdateStats()
}

func (sc *StatsContainer) ToJson() []byte {
	result, err := json.Marshal(sc)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return result
}

func (sc *StatsContainer) Save() {
	json := sc.ToJson()

	filename := sc.getFilename()

	log.Printf("Saving stats information to %s", filename)

	ioutil.WriteFile(filename, json, 0644)
}

func (rs *StatsContainer) EndRun() {
	rs.RunStats.UpdateStatsContainer(rs)
	rs.RunStats.UpdateStatsContainer(GlobalStats)
	rs.RunStats.Reset()

	rs.Save()
	GlobalStats.Save()
}

func NewStatsContainer(steamId64 string) *StatsContainer {
	characters := map[string]int{}
	causesOfDeath := map[string]int{}
	mutationChoices := map[string]int{}
	weaponChoices := map[string]int{}
	crowns := map[string]int{}

	for key, _ := range Characters {
		characters[strconv.Itoa(key)] = 0
	}

	for key, _ := range Enemies {
		causesOfDeath[strconv.Itoa(key)] = 0
	}

	for key, _ := range Mutations {
		mutationChoices[strconv.Itoa(key)] = 0
	}

	for key, _ := range Weapons {
		weaponChoices[strconv.Itoa(key)] = 0
	}

	for key, _ := range Crowns {
		crowns[strconv.Itoa(key)] = 0
	}

	sc := StatsContainer{
		steamId64,
		false,
		*NewRunStats(),
		0,
		characters,
		causesOfDeath,
		mutationChoices,
		weaponChoices,
		crowns,
		map[string]int{},
		"",
		"",
		"",
		"",
		"",
		"",
	}

	sc.load()

	return &sc
}

func init() {
	GlobalStats = NewStatsContainer("global")
}
