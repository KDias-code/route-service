package models

type Schedule struct {
	Id          int    `json:"id" db:"id"`
	RouteNumber string `json:"route_number" db:"route_number"`
	StartTime   string `json:"start_time" db:"start_time"`
	EndTime     string `json:"end_time" db:"end_time"`
}
