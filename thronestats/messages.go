package thronestats

import (
	"fmt"
	"strings"
	"log"
	"encoding/json"
	"github.com/nu7hatch/gouuid"
)

var subscriptions = map[uuid.UUID]ApiSubscriber{}
var subscribed = []string{}

func isSubscribed(steamId64 string) bool {
	for _, v := range subscribed {
		if v == steamId64 {
			return true
		}
	}

	return false
}

func SendHello(uuid uuid.UUID) {
	m := MessageOut{
		"hello",
		"Hello there!",
		"Connected to ThroneStats",
		"",
	}

	SendToConnection(uuid, m.ToJson())
}

func SendGlobalStats(uuid uuid.UUID) {
	m := MessageOut{
		"globalStats",
		"",
		string(GlobalStats.ToJson()[:]),
		"",
	}

	SendToConnection(uuid, m.ToJson())
}

func onSubscribe(uuid uuid.UUID, message MessageIn) {

	Unsubscribe(uuid)

	if isSubscribed(message.SteamId64) {
		m := MessageOut{
			"message",
			"Failed to subscribe",
			fmt.Sprintf("Someone else is subscribed to that SteamID64", message.SteamId64),
			"",
		}

		SendToConnection(uuid, m.ToJson())
		return
	}

	subscribed = append(subscribed, message.SteamId64)

	subscriber := NewApiSubscriber(uuid, message.SteamId64, strings.ToLower(message.StreamKey))

	subscriptions[uuid] = subscriber

	subscriber.Start()

	m := MessageOut{
		"message",
		"Subscribed",
		fmt.Sprintf("Now tracking runs"),
		"",
	}

	SendToConnection(uuid, m.ToJson())
}

func onStatsRequest(uuid uuid.UUID, message MessageIn) {
	sc := NewStatsContainer(message.SteamId64)

	var m MessageOut

	if (sc.Loaded == false) {
		log.Printf("Couldn't find data for player %s", message.SteamId64)
		m = MessageOut{
			"message",
			"Error",
			fmt.Sprintf("Couldn't find data for player %s", message.SteamId64),
			"",
		}
	} else {
		m = MessageOut{
			"stats",
			message.SteamId64,
			string(sc.ToJson()[:]),
			"",
		}
	}

	SendToConnection(uuid, m.ToJson())
}

func Unsubscribe(uuid uuid.UUID) {
	as, ok := subscriptions[uuid]

	if ok {
		as.Stop()
		steamId64 := subscriptions[uuid].SteamId64
		delete(subscriptions, uuid)

		newSubscribed := []string{}
		for _, v := range subscribed {
			if v != steamId64 {
				newSubscribed = append(newSubscribed, v)
			}
		}
		subscribed = newSubscribed
	}
}

func HandleMessage(uuid uuid.UUID, messageContent []byte) {

	message := MessageIn{}

	json.Unmarshal(messageContent, &message)

	switch {
	default:
		log.Printf("Unknown message received of type %s from %s", message.Type, uuid.String())

	case message.Type == "subscribe":
		log.Printf("Connection %s subscribing to %s", uuid.String(), message.SteamId64)
		onSubscribe(uuid, message)

	case message.Type == "requestStats":
		log.Printf("Connection %s requested stats for %s", uuid.String(), message.SteamId64)
		onStatsRequest(uuid, message)
	}
}
