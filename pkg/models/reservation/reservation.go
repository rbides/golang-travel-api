package models

import (
	"time"

	"github.com/google/uuid"
)

type Reservation struct {
	ID         uuid.UUID `json: "id"`
	UserID     uuid.UUID `json: "name"`
	TravelID   uuid.UUID `json: "email"`
	Created_At time.Time `json: "created_at"`
	Updated_At time.Time `json: "updated_at"`
}
