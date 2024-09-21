package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `json: "id"`
	Name       string    `json: "name"`
	Email      string    `json: "email"`
	Created_At time.Time `json: "created_at"`
	Updated_At time.Time `json: "updated_at"`
}
