package models

import "time"

type Trip struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	UserID        uint      `gorm:"not null" json:"user_id"`
	DestinationID uint      `gorm:"not null" json:"destination_id"`
	Title         string    `gorm:"not null" json:"title"`
	DurationDays  int       `gorm:"not null" json:"duration_days"`
	Budget        int       `gorm:"not null" json:"budget"`
	CreatedAt     time.Time `json:"created_at"`

	User        User        `gorm:"foreignKey:UserID"`
	Destination Destination `gorm:"foreignKey:DestinationID"`
	Itineraries []Itinerary `gorm:"foreignKey:TripID"`
}
