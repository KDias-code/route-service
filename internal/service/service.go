package service

import (
	"diplom/route-service/internal/models"
	"diplom/route-service/internal/store"
)

type IService interface {
	GetStops() ([]models.Stops, error)
	GetSchedule() ([]models.Schedule, error)
	ChangeSchedule(data models.Schedule) error
	NearStop(lat, lon string) (*models.Stops, error)
}
type Service struct {
	store store.IStore
}

func NewService(store store.IStore) *Service {
	return &Service{store: store}
}
