package models

import "github.com/google/uuid"

type Shops struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
