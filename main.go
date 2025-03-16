package main

import (
	application "NoisEsub/src/Measurement/Application"
	repositories "NoisEsub/src/Measurement/Domain/Repositories"
	adapters "NoisEsub/src/Measurement/Infraestructure/Adapters"
	controller "NoisEsub/src/Measurement/Infraestructure/Controller"
	"context"
	"log"
    

      "firebase.google.com/go"
	"github.com/gin-gonic/gin"
)

func main()  {
	ctx := context.Background()
	fbApp, err := firebase.NewApp(ctx,nil)

	if err != nil {
		log.Fatalf("failed to create firebase app: %v", err)
	}
	fcmClient, err := fbApp.Messaging(ctx)
	if err != nil {
		log.Fatal("failed to create messaging client: %v", err)
	}

	FCMSeervice := adapters.NewFCMService(fcmClient)

	var EventRepository repositories.EventRepository

	 EventService := application.NewEventService(EventRepository, FCMSeervice)
     
	 WebhookController := controller.NewWebhookController(EventService)

	 router := gin.Default()

	 router.POST("/webhooks", WebhookController.Handle )
	 
	 if err := router.Run(":8080"); err !=nil {
		log.Fatalf("failed to start server: %v", err)
	 }

}