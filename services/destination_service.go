package services

import (
	"individual-project/models"
	"individual-project/repositories"
)

type DestinationService struct {
	destRepo repositories.DestinationRepository
}

func NewDestinationService(repo repositories.DestinationRepository) *DestinationService {
	return &DestinationService{repo}
}

func (s *DestinationService) CreateDestination(dest *models.Destination) error {
	return s.destRepo.Create(dest)
}

func (s *DestinationService) GetDestinations() ([]models.Destination, error) {
	return s.destRepo.FindAll()
}

func (s *DestinationService) GetDestinationByID(id uint) (*models.Destination, error) {
	return s.destRepo.FindByID(id)
}

func (s *DestinationService) UpdateDestination(
	id uint,
	name string,
	country string,
	description string,
	cost int,
) (*models.Destination, error) {

	dest, err := s.destRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	dest.Name = name
	dest.Country = country
	dest.Description = description
	dest.EstimatedDailyCost = cost

	err = s.destRepo.Update(dest)
	if err != nil {
		return nil, err
	}

	return dest, nil
}

func (s *DestinationService) DeleteDestination(id uint) error {
	_, err := s.destRepo.FindByID(id)
	if err != nil {
		return err
	}

	return s.destRepo.Delete(id)
}
