package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"log"

	"github.com/EdvardFarrow/game-generator/internal/generator"
	"github.com/EdvardFarrow/game-generator/internal/models"
	"github.com/EdvardFarrow/game-generator/internal/publisher"
)

const (
	NumWorkers = 10    
	BufferSize = 100000

	ProjectID = "soundchain-2" 
	TopicID   = "game-events-topic"
)

func main() {
	fmt.Println("Инициализация Pub/Sub клиента...")
	
	pub, err := publisher.NewPublisher(ProjectID, TopicID)
	if err != nil {
		log.Fatalf("Не удалось создать паблишер: %v", err)
	}

	fmt.Println("Запуск генератора...")
	eventsChan := make(chan models.BaseEvent, BufferSize)

	for i := 0; i < NumWorkers; i++ {
		go worker(eventsChan)
	}

	go cloudSender(eventsChan, pub)

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)
	<-stopChan

	fmt.Println("\nПолучен сигнал остановки. Доотправляем остатки в облако...")
	pub.Stop() 
	fmt.Println("Работа корректно завершена.")
}

func worker(eventsChan chan<- models.BaseEvent) {
	for {
		event := generator.GenerateRandomEvent()
		
		eventsChan <- event
		
		// Искусственная задержка
		time.Sleep(1 * time.Millisecond) 
	}
}

func cloudSender(eventsChan <-chan models.BaseEvent, pub *publisher.GameEventPublisher) {
	ticker := time.NewTicker(1 * time.Second)
	sentCount := 0

	for {
		select {
		case event := <-eventsChan:
			pub.Publish(event)
			sentCount++
		case <-ticker.C:
			fmt.Printf("Отправлено в Pub/Sub: %d событий/сек\n", sentCount)
			sentCount = 0
		}
	}
}