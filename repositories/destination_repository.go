package repositories

import (
	"individual-project/models"

	"gorm.io/gorm"
)

type DestinationRepository interface {
	Create(destination *models.Destination) error
	FindAll() ([]models.Destination, error)
	FindByID(id uint) (*models.Destination, error)
	Update(destination *models.Destination) error
	Delete(id uint) error
}

type destinationRepository struct {
	db *gorm.DB
}

func NewDestinationRepository(db *gorm.DB) DestinationRepository {
	return &destinationRepository{db}
}

func (r *destinationRepository) Create(destination *models.Destination) error {
	return r.db.Create(destination).Error
}

func (r *destinationRepository) FindAll() ([]models.Destination, error) {
	var destinations []models.Destination
	err := r.db.Find(&destinations).Error
	return destinations, err
}

func (r *destinationRepository) FindByID(id uint) (*models.Destination, error) {
	var destination models.Destination
	err := r.db.First(&destination, id).Error
	return &destination, err
}

func (r *destinationRepository) Update(destination *models.Destination) error {
	return r.db.Save(destination).Error
}

func (r *destinationRepository) Delete(id uint) error {
	return r.db.Delete(&models.Destination{}, id).Error
}
