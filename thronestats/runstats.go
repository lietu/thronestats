package thronestats

import (
	"strconv"
)

var DEFAULT_CHARACTER = -1
var DEFAULT_CROWN = -1
var DEFAULT_ULTRA = 0
var DEFAULT_ENEMY = -1

type RunStats struct {
	character    int
	causeOfDeath int
	lastCrown    int
	lastUltra	 int
	diedOnLevel  string
	weapons      []int
	mutations    []int
	crowns       []int
}

func NewRunStats() *RunStats {
	rs := RunStats{
		DEFAULT_CHARACTER,
		DEFAULT_ENEMY,
		DEFAULT_CROWN,
		DEFAULT_ULTRA,
		"",
		[]int{},
		[]int{},
		[]int{},
	}

	return &rs
}

func (rs *RunStats) WeaponPickup(weapon int) bool {
	if IsIn(rs.weapons, weapon) {
		return false
	}

	rs.weapons = append(rs.weapons, weapon)

	return true
}

func (rs *RunStats) MutationChoice(mutation int) bool {
	if IsIn(rs.mutations, mutation) {
		return false
	}

	rs.mutations = append(rs.mutations, mutation)

	return true
}

func (rs *RunStats) CrownChoice(crown int) bool {
	if crown == DEFAULT_CROWN || crown == 0 {
		return false
	}

	if IsIn(rs.crowns, crown) {
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

func (rs *RunStats) UltraChoice(ultra int) bool {
	if ultra == 0 || ultra == rs.lastUltra {
		return false
	}

	rs.lastUltra = ultra

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
	if rs.character == DEFAULT_CHARACTER {
		return
	}

	sc.Runs += 1

	_, ok := sc.DeathsByLevel[rs.diedOnLevel]

	if !ok {
		sc.DeathsByLevel[rs.diedOnLevel] = 0
	}

	sc.DeathsByLevel[rs.diedOnLevel] += 1

	if rs.causeOfDeath != DEFAULT_ENEMY {
		sc.CausesOfDeath[strconv.Itoa(rs.causeOfDeath)] += 1
	}

	sc.Characters[strconv.Itoa(rs.character)] += 1
	sc.UltraChoices[strconv.Itoa(rs.character)][strconv.Itoa(rs.lastUltra)] += 1

	for _, mutation := range rs.mutations {
		sc.MutationChoices[strconv.Itoa(mutation)] += 1
	}

	for _, weapon := range rs.weapons {
		sc.WeaponChoices[strconv.Itoa(weapon)] += 1
	}

	for _, crown := range rs.crowns {
		if crown != DEFAULT_CROWN {
			sc.CrownChoices[strconv.Itoa(crown)] += 1
		}
	}

	sc.UpdateStats()
}

func (rs *RunStats) Reset() {
	rs.character = -1
	rs.causeOfDeath = -1
	rs.diedOnLevel = ""
	rs.lastCrown = 1
	rs.lastUltra = 0
	rs.weapons = []int{}
	rs.mutations = []int{}
	rs.crowns = []int{}
}