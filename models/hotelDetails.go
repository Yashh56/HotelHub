package models

type HotelDetails struct {
	Name           string  `json:"name"`
	Location       string  `json:"location"`
	Description    string  `json:"description"`
	Rating         float64 `json:"rating"`
	TotalRooms     int     `json:"totalRooms"`
	AvailableRooms int     `json:"availableRooms"`
}
