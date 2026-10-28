package main

import "time"

type Event struct {
	ID        string     `json:"id"`
	Type      string     `json:"type"`
	Repo      Repository `json:"repo"`
	CreatedAt time.Time  `json:"created_at"`
}

type Repository struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}
