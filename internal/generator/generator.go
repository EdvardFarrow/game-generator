package generator

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/EdvardFarrow/game-generator/internal/models"
)

func init() {
	gofakeit.Seed(0) 
}

func GenerateRandomEvent() models.BaseEvent {
	event := models.BaseEvent{
		EventID:   uuid.New().String(),
		UserID:    "user_" + gofakeit.UUID(), 
		Platform:  gofakeit.RandomString([]string{"ios", "android", "web"}),
		Timestamp: time.Now().UTC(),
	}

	if gofakeit.Number(1, 100) <= 1 {
		event.Type = models.TypeErrorSimulated
		event.Payload = map[string]string{"error_code": "503", "reason": "timeout"}
		return event
	}

	eventTypes := []models.EventType{
		models.TypeSessionStart,
		models.TypeEconomySpend,
		models.TypeLevelUp,
	}
	eventType := eventTypes[gofakeit.Number(0, len(eventTypes)-1)]
	event.Type = eventType

	switch eventType {
	case models.TypeSessionStart:
		event.Payload = models.SessionPayload{
			AppVersion: gofakeit.RandomString([]string{"1.0.1", "1.0.2", "1.1.0"}),
			Country:    gofakeit.CountryAbr(), 
		}
	case models.TypeEconomySpend:
		event.Payload = models.EconomyPayload{
			Currency: gofakeit.RandomString([]string{"soft_coin", "hard_gem"}),
			Amount:   gofakeit.Number(10, 500),
			Item:     gofakeit.RandomString([]string{"energy_refill", "sword_upgrade", "loot_box"}),
		}
	case models.TypeLevelUp:
		event.Payload = struct {
			NewLevel int `json:"new_level"`
		}{
			NewLevel: gofakeit.Number(2, 50),
		}
	}

	return event
}