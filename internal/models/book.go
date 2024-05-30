package models

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	UID       uuid.UUID `json:"uid"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	Details   string    `json:"details"`
}
