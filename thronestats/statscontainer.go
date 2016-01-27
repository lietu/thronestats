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
	SteamId64              string                      `json:"-"`
	Loaded                 bool                        `json:"-"`
	Running                bool                        `json:"-"`
	RunStats               RunStats                    `json:"-"`
	Runs                   int                         `json:"runs"`

	Characters             map[string]int              `json:"characters"`
	CausesOfDeath          map[string]int              `json:"causesOfDeath"`
	MutationChoices        map[string]int              `json:"mutationChoices"`
	WeaponChoices          map[string]int              `json:"weaponChoices"`
	CrownChoices           map[string]int              `json:"crownChoices"`
	UltraChoices           map[string]map[string]int   `json:"ultraChoices"`

	DeathsByLevel          map[string]int              `json:"deathsByLevel"`

	MostPopularWeapon      string                      `json:"mostPopularWeapon"`
	MostCommonCauseOfDeath string                      `json:"mostCommonCauseOfDeath"`
	MostPopularMutation    string                      `json:"mostPopularMutation"`
	MostPopularCrown       string                      `json:"mostPopularCrown"`
	MostPopularCharacter   string                      `json:"mostPopularCharacter"`
	MostCommonDeathLevel   string                      `json:"mostCommonDeathLevel"`
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
	return sc.getRateOfTotal(runs, total)
}

func (sc *StatsContainer) getRateOfTotal(runs int, total int) string {
	return fmt.Sprintf("%.2f%%", (float64(runs) / float64(total)) * 100)
}

func (sc *StatsContainer) GetCrownRate(mutation int, extra int) string {
	return sc.getRate(sc.CrownChoices[strconv.Itoa(mutation)] + extra, extra == 1)
}

func (sc *StatsContainer) GetMutationRate(mutation int, extra int) string {
	return sc.getRate(sc.MutationChoices[strconv.Itoa(mutation)] + extra, extra == 1)
}

func (sc *StatsContainer) GetUltraRate(character int, ultra int, extra int) string {
	runs := sc.UltraChoices[strconv.Itoa(character)][strconv.Itoa(ultra)]
	runs += extra

	total := sc.Characters[strconv.Itoa(character)] + extra

	return sc.getRateOfTotal(runs, total)
}

func (sc *StatsContainer) GetCharacterRate(character int, extra int) string {
	return sc.getRate(sc.Characters[strconv.Itoa(character)] + extra, extra == 1)
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

		sc.UpdateStats()
		sc.Loaded = true
	} else {
		if os.IsNotExist(err) {
			log.Printf("%s does not exist, do not have existing data.", filename)
		} else {
			log.Fatalf("error: %v", err)
		}
	}
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
	ultras := map[string]map[string]int{}

	for key, _ := range Characters {
		skey := strconv.Itoa(key)
		characters[skey] = 0

		count := 2
		if key == 11 {
			count = 3
		}

		ultras[skey] = map[string]int{}
		for i := 0; i <= count; i++ {
			ultras[skey][strconv.Itoa(i)] = 0
		}
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
		false,
		*NewRunStats(),
		0,
		characters,
		causesOfDeath,
		mutationChoices,
		weaponChoices,
		crowns,
		ultras,
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
