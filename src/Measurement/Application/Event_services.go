package application

import (
	entities "NoisEsub/src/Measurement/Domain/Entities"
	repositories "NoisEsub/src/Measurement/Domain/Repositories"
	"fmt"
)

type EventServices struct {
	repo repositories.EventRepository
	nifier repositories.NotificationService
}

func (h *EventServices) ProcessEvent(event entities.Event) error {
	if err := h.repo.Save(&event); err != nil {
		return fmt.Errorf("failed to save event: %w", err)
		
	}
    
	switch event.Type{

	case entities.NoiseAlert:
		if event.Data.Value > 60 {
			return h.nifier.Send("Alerta de Ruido", 
				fmt.Sprintf("Nivel: %.1f dB", event.Data.Value))
		}
	case entities.AirQualityAlert:
		if event.Data.Value > 300 {
			return h.nifier.Send("Calidad de Aire", 
				fmt.Sprintf("Nivel: %.0f ppm", event.Data.Value))
		}
	case entities.LightAlert:
		if event.Data.Value < 50 || event.Data.Value > 100 {
			return h.nifier.Send("Nivel de Luz", 
				fmt.Sprintf("Valor: %.0f lux", event.Data.Value))
		}
	}
	return nil
}