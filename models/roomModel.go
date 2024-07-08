package models

type RoomModel struct {
	RoomNumber   string  `json:"roomNumber"`
	Type         string  `json:"type"`
	Price        float64 `json:"price"`
	Availability bool    `json:"availability":`
	Description  string  `json:"description"`
}
