package services

import (
	"individual-project/models"
	"individual-project/repositories"
)

type ItineraryService struct {
	repo repositories.ItineraryRepository
}

func NewItineraryService(repo repositories.ItineraryRepository) *ItineraryService {
	return &ItineraryService{repo}
}

func (s *ItineraryService) CreateItinerary(itinerary *models.Itinerary) error {
	return s.repo.Create(itinerary)
}

func (s *ItineraryService) GetItineraries() ([]models.Itinerary, error) {
	return s.repo.FindAll()
}

func (s *ItineraryService) GetItineraryByID(id uint) (*models.Itinerary, error) {
	return s.repo.FindByID(id)
}

// UpdateItinerary updates an existing itinerary by ID
func (s *ItineraryService) UpdateItinerary(
	id uint,
	dayNumber int,
	activity string,
	location string,
	estimatedCost int,
) (*models.Itinerary, error) {

	itinerary, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	itinerary.DayNumber = dayNumber
	itinerary.Activity = activity
	itinerary.Location = location
	itinerary.EstimatedCost = estimatedCost

	err = s.repo.Update(itinerary)
	if err != nil {
		return nil, err
	}

	return itinerary, nil
}

func (s *ItineraryService) DeleteItinerary(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
