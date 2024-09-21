package models

import "github.com/google/uuid"

type SeatStatus string

const (
	Free       = SeatStatus("FREE")
	Processing = SeatStatus("PROCESSING")
	Reserved   = SeatStatus("RESERVED")
)

type Seat struct {
	ID       uuid.UUID
	TravelID uuid.UUID
	Position uint
	Status   SeatStatus
}
