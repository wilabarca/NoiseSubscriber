package entities

import "time"

type EventType string

const (
	NoiseAlert      EventType = "noise_alert"
	AirQualityAlert EventType = "air_quality_alert"
	LightAlert      EventType = "light_alert"
)

type SensorData struct {
	Value     float64   `json:"value"`
	DeviceID  string    `json:"device_id"`
	Timestamp time.Time `json:"timestamp"`
}

type Event struct {
	Type EventType `json:"type"`
	Data SensorData `json:"data"`
}