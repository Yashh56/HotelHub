package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Yashh56/HotelHub/prisma/db"
	"github.com/rs/zerolog/log"
)

type HotelDetails struct {
	Name           string  `json:"name"`
	Location       string  `json:"location"`
	Description    string  `json:"description"`
	Rating         float64 `json:"rating"`
	TotalRooms     int     `json:"totalRooms"`
	AvailableRooms int     `json:"availableRooms"`
}

// type HotelController struct {
// 	Client *db.PrismaClient
// }

func CreateHotel(client *db.PrismaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var details HotelDetails
		err := json.NewDecoder(r.Body).Decode(&details)
		if err != nil {
			http.Error(w, "Invalid Input", http.StatusBadRequest)
			log.Error().Err(err).Msg("Failed to decode request body")
			return
		}

		hotel, err := client.Hotel.CreateOne(
			db.Hotel.Name.Set(details.Name),
			db.Hotel.Location.Set(details.Location),
			db.Hotel.Description.Set(details.Description),
			db.Hotel.Rating.Set(details.Rating),
			db.Hotel.TotalRooms.Set(details.TotalRooms),
			db.Hotel.AvailableRooms.Set(details.AvailableRooms),
		).Exec(r.Context())

		if err != nil {
			http.Error(w, "Error Creating Hotel", http.StatusInternalServerError)
			log.Error().Err(err).Msg("Failed to create a hotel in database")
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(hotel)
		log.Info().Msg("Hotel Created Successfully !")
	}
}
