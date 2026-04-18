package publisher

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/EdvardFarrow/game-generator/internal/models" 
)

type GameEventPublisher struct {
	client *pubsub.Client
	topic  *pubsub.Topic
}

func NewPublisher(projectID, topicID string) (*GameEventPublisher, error) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания клиента Pub/Sub: %v", err)
	}

	topic := client.Topic(topicID)

	topic.PublishSettings.ByteThreshold = 1e6              // 1MB
	topic.PublishSettings.CountThreshold = 1000            
	topic.PublishSettings.DelayThreshold = 50 * time.Millisecond 

	return &GameEventPublisher{
		client: client,
		topic:  topic,
	}, nil
}

func (p *GameEventPublisher) Publish(event models.BaseEvent) {
	ctx := context.Background()

	data, err := json.Marshal(event)
	if err != nil {
		log.Printf("Ошибка маршалинга: %v", err)
		return
	}

	result := p.topic.Publish(ctx, &pubsub.Message{
		Data: data,
	})

	go func() {
		_, err := result.Get(ctx)
		if err != nil {
			log.Printf("Ошибка отправки в Pub/Sub: %v", err)
		}
	}()
}

func (p *GameEventPublisher) Stop() {
	p.topic.Stop()
	p.client.Close()
}