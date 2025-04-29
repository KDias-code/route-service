package service

import (
	"diplom/route-service/internal/models"
	"strconv"
)

func (s *Service) GetStops() ([]models.Stops, error) {
	return s.store.GetStops()
}

func (s *Service) NearStop(lat, lon string) (*models.Stops, error) {
	decLat, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		return nil, err
	}

	decLon, err := strconv.ParseFloat(lon, 64)
	if err != nil {
		return nil, err
	}

	resp, err := s.store.FindNearestStop(decLat, decLon)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
