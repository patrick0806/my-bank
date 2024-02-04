package dtos

import "time"

type HttpErrorDTO struct {
	StatusCode int       `json:"statusCode"`
	Reason     string    `json:"reason"`
	Timestamp  time.Time `json:"timestamp"`
}
