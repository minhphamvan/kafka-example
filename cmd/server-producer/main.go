package main

import (
	"fmt"
	"log"

	"kafka-example/env/kafka"
	"kafka-example/internal/app-api/handler/notification"
	"kafka-example/internal/pkg/kafka-utils"
	"kafka-example/internal/service-notification"
	"kafka-example/internal/store/notification"
	"kafka-example/internal/store/user"

	"github.com/gin-gonic/gin"
)

func main() {
	// New store
	producer, err := kafkautils.SetupProducer(kafkaconst.KafkaServerAddress)
	if err != nil {
		log.Fatalf("failed to initialize producer: %v", err)
	}
	defer producer.Close()

	userStore := userstore.New()
	notificationStore := notificationstore.New()

	// New service
	notificationService := servicenotification.NewService(userStore, notificationStore, producer)

	// New handler
	notificationHandler := notificationhandler.New(notificationService)

	// Setup routes
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	g := router.Group("/notifications")
	{
		g.POST("/send", notificationHandler.Send())
	}

	// Server running
	fmt.Printf("Kafka PRODUCER started at http://localhost%s\n", kafkaconst.ProducerPort)
	if err := router.Run(kafkaconst.ProducerPort); err != nil {
		log.Printf("failed to run the server: %v", err)
	}
}
