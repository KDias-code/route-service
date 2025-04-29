package service

import "diplom/route-service/internal/models"

func (s *Service) GetSchedule() ([]models.Schedule, error) {
	return s.store.GetSchedule()
}

func (s *Service) ChangeSchedule(data models.Schedule) error {
	err := s.store.ChangeSchedule(data)
	if err != nil {
		return err
	}

	return nil
}
