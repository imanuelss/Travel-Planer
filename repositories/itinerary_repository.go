package repositories

import (
	"individual-project/models"

	"gorm.io/gorm"
)

type ItineraryRepository interface {
	Create(itinerary *models.Itinerary) error
	FindAll() ([]models.Itinerary, error)
	FindByID(id uint) (*models.Itinerary, error)
	Update(itinerary *models.Itinerary) error
	Delete(id uint) error
}

type itineraryRepository struct {
	db *gorm.DB
}

func NewItineraryRepository(db *gorm.DB) ItineraryRepository {
	return &itineraryRepository{db}
}

func (r *itineraryRepository) Create(itinerary *models.Itinerary) error {
	return r.db.Create(itinerary).Error
}

func (r *itineraryRepository) FindAll() ([]models.Itinerary, error) {
	var itineraries []models.Itinerary
	err := r.db.Find(&itineraries).Error
	return itineraries, err
}

func (r *itineraryRepository) FindByID(id uint) (*models.Itinerary, error) {
	var itinerary models.Itinerary
	err := r.db.First(&itinerary, id).Error
	return &itinerary, err
}

func (r *itineraryRepository) Update(itinerary *models.Itinerary) error {
	return r.db.Save(itinerary).Error
}

func (r *itineraryRepository) Delete(id uint) error {
	return r.db.Delete(&models.Itinerary{}, id).Error
}
