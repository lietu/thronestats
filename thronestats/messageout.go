package thronestats

import (
	"log"
	"encoding/json"
)

type MessageOut struct {
	Type    string        `json:"type"`
	Header  string        `json:"header"`
	Content string        `json:"content"`
}

func (m MessageOut) ToJson() []byte {
	result, err := json.Marshal(&m)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return result
}

