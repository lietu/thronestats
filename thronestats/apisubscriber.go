package thronestats

import (
	"fmt"
	"errors"
	"time"
	"io/ioutil"
	"log"
	"strings"
	"net/http"
	"encoding/json"
	"github.com/nu7hatch/gouuid"
)

var CHECK_FREQUENCY time.Duration = 5 * time.Second

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

	header := fmt.Sprintf("%s", weapon)
	content := fmt.Sprintf("You pick up %s on %s of your runs. %s is picked up on %s of all runs.", weapon, rate, weapon, globalRate)
	icon := GetWeaponIcon(weaponCode)

	as.SendMessage(header, content, icon)

	log.Printf("Player %s picked up %s", as.SteamId64, weapon)
}

func (as *ApiSubscriber) onNewMutation(mutationCode int) {
	mutation := Mutations[mutationCode]
	rate := as.statsContainer.GetMutationRate(mutationCode, 1)
	globalRate := GlobalStats.GetMutationRate(mutationCode, 1)

	header := fmt.Sprintf("%s", mutation)
	content := fmt.Sprintf("You choose %s on %s of your runs. %s is chosen on %s of all runs.", mutation, rate, mutation, globalRate)
	icon := GetMutationIcon(mutationCode)

	as.SendMessage(header, content, icon)

	log.Printf("Player %s took %s", as.SteamId64, mutation)
}

func (as *ApiSubscriber) onNewUltra(character int, ultra int) {
	characterText := Characters[as.runData.Character]

	ultraText := Ultras[character][ultra]
	rate := as.statsContainer.GetUltraRate(character, ultra, 1)
	globalRate := GlobalStats.GetUltraRate(character, ultra, 1)

	header := "Level Ultra"
	content := fmt.Sprintf("You choose %s on %s of your %s runs. %s is chosen on %s of all %s runs.", ultraText, rate, characterText, ultraText, globalRate, characterText)
	icon := GetUltraIcon(character, ultra)

	as.SendMessage(header, content, icon)

	log.Printf("Player %s got %s ultra %d", as.SteamId64, characterText, ultra)
}

func (as *ApiSubscriber) onNewCrown(crownCode int) {
	crown := Crowns[crownCode]
	rate := as.statsContainer.GetMutationRate(crownCode, 1)
	globalRate := GlobalStats.GetMutationRate(crownCode, 1)

	header := fmt.Sprintf("%s", crown)
	content := fmt.Sprintf("You choose %s on %s of your runs. %s is chosen on %s of all runs.", crown, rate, crown, globalRate)
	icon := GetCrownIcon(crownCode)

	as.SendMessage(header, content, icon)

	log.Printf("Player %s chose %s", as.SteamId64, crown)
}

func (as *ApiSubscriber) onDeath(rd *RunData) {
	as.statsContainer.RunStats.Killed(rd.LastDamagedBy, rd.Level)
	as.statsContainer.EndRun()

	enemy := Enemies[rd.LastDamagedBy]
	rate := as.statsContainer.GetCauseOfDeathRate(rd.LastDamagedBy)
	globalRate := GlobalStats.GetCauseOfDeathRate(rd.LastDamagedBy)

	header := fmt.Sprintf("%s", enemy)
	content := fmt.Sprintf("You die to %s on %s of your runs. %s ends %s of all runs.", enemy, rate, enemy, globalRate)
	icon := GetEnemyIcon(rd.LastDamagedBy)

	as.SendMessage(header, content, icon)

	level := rd.Level

	rate = as.statsContainer.GetLevelDeathRate(level)
	globalRate = GlobalStats.GetLevelDeathRate(level)

	header = fmt.Sprintf("Died on %s", rd.Level)
	content = fmt.Sprintf("You die on %s on %s of your runs.  %s of all runs end at %s.", level, rate, globalRate, level)
	icon = GetCharacterIcon(rd.Character, rd.BSkin)

	as.SendMessage(header, content, icon)

	log.Printf("Player %s died to a %s on %s", as.SteamId64, enemy, level)
}

func (as *ApiSubscriber) onNewRun(rd *RunData) {
	as.statsContainer.RunStats.NewRun(rd.Character)

	character := Characters[rd.Character]
	rate := as.statsContainer.GetCharacterRate(rd.Character, 1)
	globalRate := GlobalStats.GetCharacterRate(rd.Character, 1)

	header := fmt.Sprintf("%s", character)
	content := fmt.Sprintf("You use %s on %s of your runs. %s is used on %s of all runs.", character, rate, character, globalRate)
	icon := GetCharacterIcon(rd.Character, rd.BSkin)

	as.SendMessage(header, content, icon)

	log.Printf("Player %s started a new run with %s", as.SteamId64, character)
}

func (as *ApiSubscriber) onNewLevel(rd *RunData) {
	log.Printf("Player %s entered level %s", as.SteamId64, rd.Level)
}

func (as *ApiSubscriber) processUpdate(rdc *RunDataContainer) {
	current := rdc.Current
	previous := rdc.Previous

	// Has the player died?
	if previous.Timestamp > 0 && previous.Timestamp == as.runData.Timestamp {
		as.onDeath(previous)
		as.runData.Timestamp = 0
	}

	// If we have a current run
	if rdc.Current.Timestamp > 0 {
		if as.statsContainer.Running == false {
			as.onNewRun(current)
			as.statsContainer.Running = true
		} else {
			for _, v := range current.Weapons {
				if as.statsContainer.RunStats.WeaponPickup(v) {
					as.onWeaponPickup(v)
				}
			}

			for _, v := range current.Mutations {
				if as.statsContainer.RunStats.MutationChoice(v) {
					as.onNewMutation(v)
				}
			}

			if as.statsContainer.RunStats.CrownChoice(current.Crown) {
				as.onNewCrown(current.Crown)
			}

			if as.statsContainer.RunStats.UltraChoice(current.Ultra) {
				as.onNewUltra(current.Character, current.Ultra)
			}

			// Reached new level
			if as.runData.Level != current.Level {
				as.onNewLevel(current)
			}
		}

		as.runData = rdc.Current
	} else {
		as.statsContainer.Running = false
	}
}

func (as *ApiSubscriber) getData() ([]byte, error) {
	url := fmt.Sprintf("https://tb-api.xyz/stream/get?s=%s&key=%s", as.SteamId64, strings.ToUpper(as.StreamKey))

	resp, err := http.Get(url)

	if err != nil {
		log.Printf("API error: %s", err)
		return nil, errors.New("API error")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Printf("Error reading API response: %s", err)
		return nil, errors.New("API error")
	}

	invalidSettings := false
	if resp.StatusCode != 200 {
		invalidSettings = true
	}

	if strings.Contains(string(body[:]), "<html>") {
		invalidSettings = true
	}

	if invalidSettings {
		if as.invalidSettings == false {
			msg := fmt.Sprintf("Could not get data with the provided Steam ID and Stream Key. Have you done any runs that would've gotten recorded?")
			as.SendMessage("Invalid settings?", msg, "")
			as.invalidSettings = true
		}
		return nil, errors.New("API error")
	}

	return body, nil
}

func (as *ApiSubscriber) poll() {
	body, err := as.getData()

	if err != nil {
		return
	}

	response := NewApiResponse()

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Printf("JSON error decoding API response: %s", err)
		log.Printf("%s", string(body[:]))
		return
	}

	data := response.ToRunData()
	as.processUpdate(data)
}

func (as *ApiSubscriber) run() {
	start := time.Now()

	as.poll()

	for {
		select {
		case <-as.done:
			log.Printf("ApiSubscriber for %s stopping.", as.SteamId64)
			return

		default:
			elapsed := time.Since(start)

			if elapsed > CHECK_FREQUENCY {
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
