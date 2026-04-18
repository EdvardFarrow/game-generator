package models

import "time"

type EventType string

const (
	TypeSessionStart   EventType = "session_start"
	TypeEconomySpend   EventType = "economy_spend"
	TypeLevelUp        EventType = "level_up"
	TypeErrorSimulated EventType = "simulated_error" 
)

type BaseEvent struct {
	EventID   string    `json:"event_id"`
	UserID    string    `json:"user_id"`
	Platform  string    `json:"platform"` 
	Timestamp time.Time `json:"timestamp"`
	Type      EventType `json:"type"`
	
	Payload   interface{} `json:"payload"` 
}

type SessionPayload struct {
	AppVersion string `json:"app_version"`
	Country    string `json:"country"`
}

type EconomyPayload struct {
	Currency string `json:"currency"`
	Amount   int    `json:"amount"`
	Item     string `json:"item"`     
}
