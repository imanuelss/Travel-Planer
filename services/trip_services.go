package services

import (
	"individual-project/models"
	"individual-project/repositories"
)

type TripService struct {
	tripRepo repositories.TripRepository
}

func NewTripService(repo repositories.TripRepository) *TripService {
	return &TripService{repo}
}

func (s *TripService) CreateTrip(trip *models.Trip) error {
	return s.tripRepo.Create(trip)
}

func (s *TripService) GetTrips() ([]models.Trip, error) {
	return s.tripRepo.FindAll()
}

func (s *TripService) GetTripByID(id uint) (*models.Trip, error) {
	return s.tripRepo.FindByID(id)
}

// UpdateTrip updates an existing trip by ID
func (s *TripService) UpdateTrip(id uint, title string, durationDays, budget int) (*models.Trip, error) {
	trip, err := s.tripRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	trip.Title = title
	trip.DurationDays = durationDays
	trip.Budget = budget

	err = s.tripRepo.Update(trip)
	if err != nil {
		return nil, err
	}

	return trip, nil
}

func (s *TripService) DeleteTrip(id uint) error {
	_, err := s.tripRepo.FindByID(id)
	if err != nil {
		return err
	}

	return s.tripRepo.Delete(id)
}
