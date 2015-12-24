package thronestats

import (
	"strconv"
)

type RunStats struct {
	character    int
	causeOfDeath int
	diedOnLevel  string
	lastCrown    int
	weapons      []int
	mutations    []int
	crowns       []int
}

func is_in(list []int, value int) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}

	return false
}

func (rs *RunStats) WeaponPickup(weapon int) bool {
	if is_in(rs.weapons, weapon) {
		return false
	}

	rs.weapons = append(rs.weapons, weapon)

	return true
}

func (rs *RunStats) MutationChoice(mutation int) bool {
	if is_in(rs.mutations, mutation) {
		return false
	}

	rs.mutations = append(rs.mutations, mutation)

	return true
}

func (rs *RunStats) CrownChoice(crown int) bool {
	if is_in(rs.crowns, crown) {
		if crown != rs.lastCrown {
			rs.lastCrown = crown
			return true
		}
		return false
	}

	rs.crowns = append(rs.crowns, crown)

	rs.lastCrown = crown

	// Skip notification for starting with "Bare head"
	if crown == 1 {
		return false
	}

	return true
}

func (rs *RunStats) Killed(causeOfDeath int, diedOnLevel string) {
	rs.causeOfDeath = causeOfDeath
	rs.diedOnLevel = diedOnLevel
}

func (rs *RunStats) NewRun(character int) {
	rs.character = character
}

func (rs *RunStats) UpdateStatsContainer(sc *StatsContainer) {
	_, ok := sc.DeathsByLevel[rs.diedOnLevel]

	if !ok {
		sc.DeathsByLevel[rs.diedOnLevel] = 0
	}

	sc.DeathsByLevel[rs.diedOnLevel] += 1
	sc.CausesOfDeath[strconv.Itoa(rs.causeOfDeath)] += 1
	sc.Characters[strconv.Itoa(rs.character)] += 1

	for _, mutation := range rs.mutations {
		sc.MutationChoices[strconv.Itoa(mutation)] += 1
	}

	for _, weapon := range rs.weapons {
		sc.WeaponChoices[strconv.Itoa(weapon)] += 1
	}

	for _, crown := range rs.crowns {
		sc.CrownChoices[strconv.Itoa(crown)] += 1
	}

	// Calculate stats
	weapon := sc.getMostPopular(sc.WeaponChoices, "1")
	causeOfDeath := sc.getMostPopular(sc.CausesOfDeath, "")
	mutation := sc.getMostPopular(sc.MutationChoices, "")
	crown := sc.getMostPopular(sc.CrownChoices, "1")
	character := sc.getMostPopular(sc.Characters, "")
	level := sc.getMostPopular(sc.DeathsByLevel, "")

	sc.MostPopularWeapon = Weapons[ToInt(weapon)]
	sc.MostCommonCauseOfDeath = Enemies[ToInt(causeOfDeath)]
	sc.MostPopularMutation = Mutations[ToInt(mutation)]
	sc.MostPopularCrown = Crowns[ToInt(crown)]
	sc.MostPopularCharacter = Characters[ToInt(character)]
	sc.MostCommonDeathLevel = level
}

func (rs *RunStats) Reset() {
	rs.character = -1
	rs.causeOfDeath = -1
	rs.diedOnLevel = ""
	rs.lastCrown = 1
	rs.weapons = []int{}
	rs.mutations = []int{}
	rs.crowns = []int{}
}

func NewRunStats() *RunStats {
	rs := RunStats{
		-1,
		-1,
		"",
		1,
		[]int{},
		[]int{},
		[]int{},
	}

	return &rs
}