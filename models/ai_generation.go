package models

import "time"

type AIGeneration struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	TripID          uint      `gorm:"not null" json:"trip_id"`
	Prompt          string    `gorm:"type:text;not null" json:"prompt"`
	GeneratedResult string    `gorm:"type:text;not null" json:"generated_result"`
	CreatedAt       time.Time `json:"created_at"`

	Trip Trip `gorm:"foreignKey:TripID"`
}
