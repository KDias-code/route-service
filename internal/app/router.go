package app

func (s *server) router() {
	s.app.Get("/get-stops", s.handlers.GetStops)
	s.app.Get("/get-schedule", s.handlers.GetSchedule)
	s.app.Put("/update-schedule", s.handlers.UpdateSchedule)
	s.app.Get("/near-stop", s.handlers.NearStop)
}
