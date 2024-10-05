package models

import "time"

type History struct {
	ID         int       `json:"id"`
	CustomerID int       `json:"customer_id"`
	Action     string    `json:"action"`
	Timestamp  time.Time `json:"timestamp"`
}
