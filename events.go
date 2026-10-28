package main

import (
	"encoding/json"
	"time"
)

type Event struct {
	ID        string          `json:"id"`
	Type      string          `json:"type"`
	Repo      Repository      `json:"repo"`
	Payload   json.RawMessage `json:"payload"`
	CreatedAt time.Time       `json:"created_at"`
}

type PushEventPayload struct {
	Size int    `json:"size"`
	Ref  string `json:"ref"`
}

type ActionPayload struct {
	Action string `json:"action"`
}

type Repository struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}
