package repository

import "time"

type Report struct {
	ID         uint      `json:"id"`
	UserID     uint      `json:"user_id gorm:column:user_id"`
	PinpointID uint      `json:"pinpoint_id gorm:column:pinpoint_id"`
	Reason     string    `json:"reason"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
