package adapters

import (
	
"context"
	"log"
	"firebase.google.com/go/messaging"
)


type FCMService struct{
	client *messaging.Client
}

 func NewFCMService(client *messaging.Client) *FCMService {
	 return &FCMService{client: client}
	
 }

 func (s *FCMService) Send(title, body string) error {
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Topic: "alerts",
	}

	_, err := s.client.Send(context.Background(), message)
	if err != nil {
		log.Printf("FCM error: %v", err)
	}
	return err
}