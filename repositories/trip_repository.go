package repositories

import (
	"individual-project/models"

	"gorm.io/gorm"
)

type TripRepository interface {
	Create(trip *models.Trip) error
	FindAll() ([]models.Trip, error)
	FindByID(id uint) (*models.Trip, error)
	Update(trip *models.Trip) error
	Delete(id uint) error
}

type tripRepository struct {
	db *gorm.DB
}

func NewTripRepository(db *gorm.DB) TripRepository {
	return &tripRepository{db}
}

func (r *tripRepository) Create(trip *models.Trip) error {
	return r.db.Create(trip).Error
}

func (r *tripRepository) FindAll() ([]models.Trip, error) {
	var trips []models.Trip
	err := r.db.Preload("Destination").Find(&trips).Error
	return trips, err
}

func (r *tripRepository) FindByID(id uint) (*models.Trip, error) {
	var trip models.Trip
	err := r.db.Preload("Destination").First(&trip, id).Error
	return &trip, err
}

func (r *tripRepository) Update(trip *models.Trip) error {
	return r.db.Save(trip).Error
}

func (r *tripRepository) Delete(id uint) error {
	return r.db.Delete(&models.Trip{}, id).Error
}
