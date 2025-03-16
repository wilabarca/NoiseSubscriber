package application

import (
	entities "NoisEsub/src/Measurement/Domain/Entities"
	repositories "NoisEsub/src/Measurement/Domain/Repositories"
	"fmt"
)

type EventService struct {
	repo    repositories.EventRepository
	notifier repositories.NotificationService
}

func NewEventService(repo repositories.EventRepository, notifier repositories.NotificationService) *EventService {
	return &EventService{
		repo:     repo,
		notifier: notifier,
	}
}

func (s *EventService) ProcessEvent(event entities.Event) error {
	if err := s.repo.Save(&event); err != nil {
		return fmt.Errorf("failed to save event: %w", err)
	}

	switch event.Type {
	case entities.NoiseAlert:
		if event.Data.Value > 60 {
			return s.notifier.Send("Alerta de Ruido", 
				fmt.Sprintf("Nivel: %.1f dB", event.Data.Value))
		}
	case entities.AirQualityAlert:
		if event.Data.Value > 300 {
			return s.notifier.Send("Calidad de Aire", 
				fmt.Sprintf("Nivel: %.0f ppm", event.Data.Value))
		}
	case entities.LightAlert:
		if event.Data.Value < 50 || event.Data.Value > 100 {
			return s.notifier.Send("Nivel de Luz", 
				fmt.Sprintf("Valor: %.0f lux", event.Data.Value))
		}
	}
	return nil
}