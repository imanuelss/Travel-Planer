package services

import (
	"individual-project/models"
	"individual-project/repositories"
)

type AIGenerationService struct {
	repo repositories.AIGenerationRepository
}

func NewAIGenerationService(repo repositories.AIGenerationRepository) *AIGenerationService {
	return &AIGenerationService{
		repo: repo,
	}
}

func (s *AIGenerationService) Create(ai *models.AIGeneration) error {
	return s.repo.Create(ai)
}

func (s *AIGenerationService) GetAll() ([]models.AIGeneration, error) {
	return s.repo.FindAll()
}

func (s *AIGenerationService) GetByID(id uint) (*models.AIGeneration, error) {
	return s.repo.FindByID(id)
}
