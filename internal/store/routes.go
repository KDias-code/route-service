package store

import (
	"diplom/route-service/internal/models"
)

func (s *Store) GetStops() ([]models.Stops, error) {
	var stops []models.Stops

	err := s.db.Select(&stops, "SELECT * FROM stops")
	if err != nil {
		return stops, err
	}

	return stops, nil
}

func (s *Store) GetSchedule() ([]models.Schedule, error) {
	var schedule []models.Schedule

	err := s.db.Select(&schedule, "SELECT * FROM routes")
	if err != nil {
		return schedule, err
	}

	return schedule, nil
}

func (s *Store) ChangeSchedule(data models.Schedule) error {
	query := "UPDATE routes SET route_number = $1, start_time = $2, end_time = $3 WHERE id = $4"

	_, err := s.db.Exec(query, data.RouteNumber, data.StartTime, data.EndTime, data.Id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) FindNearestStop(lat, lon float64) (*models.Stops, error) {
	query := `
        SELECT id, stop_name, latitude, longitude, route_id,
        ( 6371 * acos( cos( radians($1) ) * cos( radians(latitude) ) * cos( radians(longitude) - radians($2) ) + sin( radians($1) ) * sin( radians(latitude) ) ) ) AS distance
        FROM stops
        ORDER BY distance
        LIMIT 1;
    `

	row := s.db.QueryRow(query, lat, lon)
	var stop models.Stops
	var distance float64

	err := row.Scan(&stop.Id, &stop.StopName, &stop.Latitude, &stop.Longitude, &stop.RouteId, &distance)
	if err != nil {
		return nil, err
	}

	return &stop, nil
}
