package thronestats

import (
	"fmt"
	"time"
	"io/ioutil"
	"log"
	"strings"
	"net/http"
	"encoding/json"
	"github.com/nu7hatch/gouuid"
)

type ApiSubscriber struct {
	ClientUUID      uuid.UUID
	SteamId64       string
	StreamKey       string
	done            chan bool
	runData         *RunData
	statsContainer  *StatsContainer
	invalidSettings bool
}

func (as *ApiSubscriber) onWeaponPickup(weaponCode int) {
	weapon := Weapons[weaponCode]
	rate := as.statsContainer.GetWeaponRate(weaponCode, 1)
	globalRate := GlobalStats.GetWeaponRate(weaponCode, 1)

	header := fmt.Sprintf("%s!", weapon)
	content := fmt.Sprintf("You pick up %s on %s of your runs. %s is picked up on %s of all runs.", weapon, rate, weapon, globalRate)
	icon := ""

	as.SendMessage(header, content, icon)

	log.Printf("Player %s picked up %s", as.SteamId64, weapon)
}

func (as *ApiSubscriber) onNewMutation(mutationCode int) {
	mutation := Mutations[mutationCode]
	rate := as.statsContainer.GetMutationRate(mutationCode, 1)
	globalRate := GlobalStats.GetMutationRate(mutationCode, 1)

	header := fmt.Sprintf("%s!", mutation)
	content := fmt.Sprintf("You choose %s on %s of your runs. %s is chosen on %s of all runs.", mutation, rate, mutation, globalRate)
	icon := GetMutationIcon(mutationCode)

	as.SendMessage(header, content, icon)

	log.Printf("Player %s took %s", as.SteamId64, mutation)
}

func (as *ApiSubscriber) onNewUltra(character int, ultra int) {
	characterText := Characters[as.runData.Character]
	log.Printf("Player %s got %s ultra %d", as.SteamId64, characterText, ultra)
}

func (as *ApiSubscriber) onNewCrown(crownCode int) {
	crown := Crowns[crownCode]
	rate := as.statsContainer.GetMutationRate(crownCode, 1)
	globalRate := GlobalStats.GetMutationRate(crownCode, 1)

	header := fmt.Sprintf("%s!", crown)
	content := fmt.Sprintf("You choose %s on %s of your runs. %s is chosen on %s of all runs.", crown, rate, crown, globalRate)
	icon := ""

	as.SendMessage(header, content, icon)

	log.Printf("Player %s chose %s", as.SteamId64, crown)
}

func (as *ApiSubscriber) onDeath() {
	as.statsContainer.RunStats.Killed(as.runData.LastDamagedBy, as.runData.Level)
	as.statsContainer.EndRun()


	enemy := Enemies[as.runData.LastDamagedBy]
	rate := as.statsContainer.GetCauseOfDeathRate(as.runData.LastDamagedBy)
	globalRate := GlobalStats.GetCauseOfDeathRate(as.runData.LastDamagedBy)

	header := fmt.Sprintf("%s", enemy)
	content := fmt.Sprintf("You die to %s on %s of your runs. %s ends %s of all runs.", enemy, rate, enemy, globalRate)
	icon := ""

	as.SendMessage(header, content, icon)


	level := as.runData.Level

	rate = as.statsContainer.GetLevelDeathRate(level)
	globalRate = GlobalStats.GetLevelDeathRate(level)

	header = fmt.Sprintf("Died on %s", as.runData.Level)
	content = fmt.Sprintf("You die on %s on %s of your runs.  %s of all runs end at %s.", level, rate, globalRate, level)
	icon = GetCharacterIcon(as.runData.Character, as.runData.BSkin)

	as.SendMessage(header, content, icon)

	log.Printf("Player %s died to a %s on %s", as.SteamId64, enemy, level)
}

func (as *ApiSubscriber) onNewRun(rd *RunData) {
	as.statsContainer.RunStats.NewRun(rd.Character)

	character := Characters[rd.Character]
	rate := as.statsContainer.GetCharacterRate(rd.Character, 1)
	globalRate := GlobalStats.GetCharacterRate(rd.Character, 1)

	header := fmt.Sprintf("New run with %s", character)
	content := fmt.Sprintf("You use %s on %s of your runs. %s is used on %s of all runs.", character, rate, character, globalRate)
	icon := GetCharacterIcon(rd.Character, rd.BSkin)

	as.SendMessage(header, content, icon)

	log.Printf("Player %s started a new run with %s", as.SteamId64, character)
}

func (as *ApiSubscriber) onNewLevel(rd *RunData) {
	log.Printf("Player %s entered level %s", as.SteamId64, rd.Level)
}

func (as *ApiSubscriber) processUpdate(rd *RunData) {
	if as.statsContainer.Running {
		for _, v := range rd.Weapons {
			if as.statsContainer.RunStats.WeaponPickup(v) {
				as.onWeaponPickup(v)
			}
		}

		for _, v := range rd.Mutations {
			if as.statsContainer.RunStats.MutationChoice(v) {
				as.onNewMutation(v)
			}
		}

		if as.statsContainer.RunStats.CrownChoice(rd.Crown) {
			as.onNewCrown(rd.Crown)
		}


		// Reached new level
		if as.runData.Level != rd.Level {
			as.onNewLevel(rd)
		}

		// Has the player died?
		if as.runData.RunTime > rd.RunTime {
			as.onDeath()
			as.onNewRun(rd)
		}
	}

	as.statsContainer.Running = true
	as.runData = rd
}

func (as *ApiSubscriber) poll() {
	url := fmt.Sprintf("http://nuclearthrone.com/data/players/data/%s%s.cur", as.SteamId64, as.StreamKey)

	resp, err := http.Get(url)

	if err != nil {
		log.Printf("API error: %s", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Printf("Error reading API response: %s", err)
		return
	}

	if strings.Contains(string(body[:]), "<html>") {
		if as.invalidSettings == false {
			msg := fmt.Sprintf("Could not get data with the provided Steam ID and Stream Key. Have you done any runs that would've gotten recorded?")
			as.SendMessage("Invalid settings?", msg, "")
			as.invalidSettings = true
		}
		return
	}

	response := ApiResponse{}

	json.Unmarshal(body, &response)

	data := response.ToRunData()
	as.processUpdate(&data)
}

func (as *ApiSubscriber) run() {
	checkFrequency := 5 * time.Second
	start := time.Now()

	as.poll()

	for {
		select {
		case <-as.done:
			log.Printf("ApiSubscriber for %s stopping.", as.SteamId64)
			return

		default:
			elapsed := time.Since(start)

			if elapsed > checkFrequency {
				start = time.Now()
				as.poll()
			} else {
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}

func (as *ApiSubscriber) Start() {
	go as.run()
}

func (as *ApiSubscriber) Stop() {
	log.Printf("Asking ApiSubscriber for %s to stop.", as.SteamId64)
	as.done <- true
}

func (as *ApiSubscriber) SendMessage(header string, content string, icon string) {
	m := MessageOut{
		"message",
		header,
		content,
		icon,
	}

	SendToConnection(as.ClientUUID, m.ToJson())
}

func NewApiSubscriber(uuid uuid.UUID, steamId64 string, streamKey string) ApiSubscriber {
	as := ApiSubscriber{
		uuid,
		steamId64,
		streamKey,
		make(chan bool),
		NewRunData(),
		NewStatsContainer(steamId64),
		false,
	}

	return as
}
