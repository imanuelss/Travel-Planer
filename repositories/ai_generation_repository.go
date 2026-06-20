package repositories

import (
	"individual-project/models"

	"gorm.io/gorm"
)

type AIGenerationRepository interface {
	Create(ai *models.AIGeneration) error
	FindAll() ([]models.AIGeneration, error)
	FindByID(id uint) (*models.AIGeneration, error)
}

type aiGenerationRepository struct {
	db *gorm.DB
}

func NewAIGenerationRepository(db *gorm.DB) AIGenerationRepository {
	return &aiGenerationRepository{db}
}

func (r *aiGenerationRepository) Create(ai *models.AIGeneration) error {
	return r.db.Create(ai).Error
}

func (r *aiGenerationRepository) FindAll() ([]models.AIGeneration, error) {
	var data []models.AIGeneration
	err := r.db.Find(&data).Error
	return data, err
}

func (r *aiGenerationRepository) FindByID(id uint) (*models.AIGeneration, error) {
	var data models.AIGeneration
	err := r.db.First(&data, id).Error
	return &data, err
}
