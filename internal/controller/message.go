package controller

import "encoding/json"

type Message struct {
	Action string          `json:"action"`
	Params json.RawMessage `json:"params"`
}
