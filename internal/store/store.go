package store

import (
	"diplom/route-service/internal/models"
	"github.com/jmoiron/sqlx"
)

type IStore interface {
	GetStops() ([]models.Stops, error)
	GetSchedule() ([]models.Schedule, error)
	ChangeSchedule(data models.Schedule) error
	FindNearestStop(lat, lon float64) (*models.Stops, error)
}

type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{db: db}
}
