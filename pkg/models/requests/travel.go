package request

import (
	"time"
)

type TravelRequest struct {
	Name        string    `json:"name" binding:"required"`
	Destination string    `json:"destination" binding:"required"`
	Price       float64   `json:"price" binding:"required"`
	Seats       uint      `json:"seats" binding:"required"`
	Departure   time.Time `json:"departure" binding:"required"`
}
