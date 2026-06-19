package models

import "time"

type Destination struct {
	ID                 uint      `gorm:"primaryKey" json:"id"`
	Name               string    `gorm:"not null" json:"name"`
	Country            string    `gorm:"not null" json:"country"`
	Description        string    `json:"description"`
	EstimatedDailyCost int       `json:"estimated_daily_cost"`
	CreatedAt          time.Time `json:"created_at"`

	Trips []Trip `gorm:"foreignKey:DestinationID"`
}
