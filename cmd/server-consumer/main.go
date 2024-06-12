package main

import (
	"context"
	"fmt"
	"log"

	"kafka-example/env/kafka"
	"kafka-example/internal/app-api/handler/notification"
	"kafka-example/internal/service-notification"
	"kafka-example/internal/store/notification"
	"kafka-example/internal/store/user"

	"github.com/gin-gonic/gin"
)

func main() {
	// New store
	userRepo := userstore.New()
	notificationRepo := notificationstore.New()

	// New service
	notificationService := servicenotification.NewService(userRepo, notificationRepo, nil)

	// New handler
	notificationHandler := notificationhandler.New(notificationService)

	ctx, cancel := context.WithCancel(context.Background())
	go setupConsumerGroup(ctx, notificationRepo)
	defer cancel()

	// Setup routes
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	g := router.Group("/notifications")
	{
		g.GET("/:user_id", notificationHandler.GetByUserID())
	}

	// Server running
	fmt.Printf("Kafka CONSUMER (Group: %s) started at http://localhost%s\n",
		kafkaconst.ConsumerGroupNotifications, kafkaconst.ConsumerPort)

	if err := router.Run(kafkaconst.ConsumerPort); err != nil {
		log.Printf("failed to run the server: %v", err)
	}
}
