package repositories

import entities "NoisEsub/src/Measurement/Domain/Entities"

type EventRepository interface {
	Save(event *entities.Event) error
}

type NotificationService interface {
	Send(title, body string) error
}