package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Yashh56/HotelHub/models"
	"github.com/Yashh56/HotelHub/prisma/db"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

// type HotelDetails struct {
// 	Name           string  `json:"name"`
// 	Location       string  `json:"location"`
// 	Description    string  `json:"description"`
// 	Rating         float64 `json:"rating"`
// 	TotalRooms     int     `json:"totalRooms"`
// 	AvailableRooms int     `json:"availableRooms"`
// }

func CreateHotel(client *db.PrismaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var details models.HotelDetails
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

func GetAllHotels(client *db.PrismaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hotels, err := client.Hotel.FindMany().Exec(r.Context())
		if err != nil {
			http.Error(w, "Error fetching hotels", http.StatusInternalServerError)
			log.Error().Err(err).Msg("Failed to fetch hotels from database")
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(hotels)
	}
}

func GetHotelById(client *db.PrismaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]

		hotel, err := client.Hotel.FindUnique(
			db.Hotel.ID.Equals(id),
		).Exec(r.Context())

		if err != nil {
			http.Error(w, "Error fetching hotels", http.StatusInternalServerError)
			log.Error().Err(err).Msg("Failed to fetch hotels from database")
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(hotel)
	}
}

func UpdateHotel(client *db.PrismaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]

		var input models.HotelDetails

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "error in updating", http.StatusBadRequest)
		}

		if id == "" {
			http.Error(w, "hotel not exists", http.StatusBadRequest)
			return
		}

		updatedHotel, err := client.Hotel.FindUnique(
			db.Hotel.ID.Equals(id),
		).Update(
			db.Hotel.Name.Set(input.Name),
			db.Hotel.Location.Set(input.Location),
			db.Hotel.Description.Set(input.Description),
			db.Hotel.Rating.Set(input.Rating),
			db.Hotel.TotalRooms.Set(input.TotalRooms),
			db.Hotel.AvailableRooms.Set(input.AvailableRooms),
		).Exec(r.Context())

		if err != nil {
			log.Error().Err(err).Msg("Failed to update hotel")
			http.Error(w, "Failed to update hotel", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(updatedHotel)
	}
}

func DeleteHotel(client *db.PrismaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]

		delete, err := client.Hotel.FindMany(
			db.Hotel.ID.Equals(id),
		).Delete().Exec(r.Context())

		if err != nil {
			http.Error(w, "Failed to delete Hotel", http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(delete)
	}
}
