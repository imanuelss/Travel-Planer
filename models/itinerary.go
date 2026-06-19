package models

type Itinerary struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	TripID        uint   `gorm:"not null" json:"trip_id"`
	DayNumber     int    `gorm:"not null" json:"day_number"`
	Activity      string `gorm:"not null" json:"activity"`
	Location      string `json:"location"`
	EstimatedCost int    `json:"estimated_cost"`

	Trip Trip `gorm:"foreignKey:TripID"`
}
