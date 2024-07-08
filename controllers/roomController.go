package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Yashh56/HotelHub/models"
	"github.com/Yashh56/HotelHub/prisma/db"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func CreateRoom(client *db.PrismaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		hotelID := vars["id"]

		if hotelID == "" {
			http.Error(w, "Hotel ID is required", http.StatusBadRequest)
			log.Error().Msg("Hotel ID is missing in URL parameters")
			return
		}

		var newRoom models.RoomModel
		err := json.NewDecoder(r.Body).Decode(&newRoom)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			log.Error().Err(err).Msg("Failed to decode request body")
			return
		}

		createdRoom, err := client.Room.CreateOne(
			db.Room.RoomNumber.Set(newRoom.RoomNumber),
			db.Room.Type.Set(newRoom.Type),
			db.Room.Price.Set(newRoom.Price),
			db.Room.Description.Set(newRoom.Description),
			db.Room.Hotel.Link(db.Hotel.ID.Equals(hotelID)),
			db.Room.CreatedAt.Set(time.Now()),
			db.Room.UpdatedAt.Set(time.Now()),
		).Exec(r.Context())
		if err != nil {
			http.Error(w, "Error creating room", http.StatusInternalServerError)
			log.Error().Err(err).Msg("Failed to create room in database")
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdRoom)
		log.Info().Msg("Room created successfully")
	}
}

func GetRooms(client *db.PrismaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rooms, err := client.Room.FindMany().Exec(r.Context())
		if err != nil {
			http.Error(w, "Error creating room", http.StatusInternalServerError)
			log.Error().Err(err).Msg("Failed to create room in database")
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(rooms)

	}
}

func GetRoomById(client *db.PrismaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]

		room, err := client.Room.FindUnique(
			db.Room.ID.Equals(id),
		).Exec(r.Context())

		if err != nil {
			http.Error(w, "Error fetching room", http.StatusInternalServerError)
			log.Error().Err(err).Msg("Failed to fetch room from database")
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(room)
	}
}
