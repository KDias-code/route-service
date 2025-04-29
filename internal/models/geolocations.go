package models

type Stops struct {
	Id        int    `json:"id" db:"id"`
	StopName  string `json:"stop_name" db:"stop_name"`
	RouteId   int    `json:"route_id" db:"route_id"`
	Latitude  string `json:"latitude" db:"latitude"`
	Longitude string `json:"longitude" db:"longitude"`
}
