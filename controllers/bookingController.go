package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Yashh56/HotelHub/models"
	"github.com/Yashh56/HotelHub/prisma/db"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func CreateBooking(client *db.PrismaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var booking models.BookingModel
		userID := r.Context().Value("userId").(string)

		params := mux.Vars(r)
		hotelId := params["hotelId"]

		err := json.NewDecoder(r.Body).Decode(&booking)
		if err != nil {
			http.Error(w, "Invalid Input", http.StatusBadRequest)
			log.Error().Err(err).Msg("Failed to decode request body")
			return
		}

		checkInTime, err := time.Parse("2006-01-02", booking.CheckInDate)
		if err != nil {
			http.Error(w, "Invalid Input", http.StatusBadRequest)
			log.Error().Err(err).Msg("Failed to parse CheckInDate")
			return
		}
		checkOutTime, err := time.Parse("2006-01-02", booking.CheckInDate)
		if err != nil {
			http.Error(w, "Invalid Input", http.StatusBadRequest)
			log.Error().Err(err).Msg("Failed to parse CheckInDate")
			return
		}

		createBooking, err := client.Booking.CreateOne(
			db.Booking.CheckInDate.Set(checkInTime),
			db.Booking.CheckOutDate.Set(checkOutTime),
			db.Booking.PaymentStatus.Set(booking.PaymentStatus),
			db.Booking.User.Link(db.User.ID.Equals(userID)),
			db.Booking.Room.Link(db.Room.ID.Equals(booking.RoomId)),
			db.Booking.Customer.Link(db.Customer.ID.Equals(booking.CustomerId)),
			db.Booking.Hotel.Link(db.Hotel.ID.Equals(hotelId)),
			db.Booking.CreatedAt.Set(time.Now()),
			db.Booking.UpdatedAt.Set(time.Now()),
		).Exec(r.Context())

		if err != nil {
			http.Error(w, "Error creating room", http.StatusInternalServerError)
			log.Error().Err(err).Msg("Failed to create booking in database")
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createBooking)
		log.Info().Msg("Booking has completed")
		fmt.Println(db.Booking.ID)
	}
}

func GetBookings(client *db.PrismaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["hotelId"]

		bookings, err := client.Booking.FindMany(
			db.Booking.ID.Equals(id),
		).Exec(r.Context())

		if err != nil {
			http.Error(w, "Error creating room", http.StatusInternalServerError)
			log.Error().Err(err).Msg("Failed to Load booking in database")
			return
		}

		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(bookings)
		log.Info().Msg("All bookings")
	}
}

func GetBookingById(client *db.PrismaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]

		byId, err := client.Booking.FindUnique(
			db.Booking.ID.Equals(id),
		).Exec(r.Context())
		if err != nil {
			http.Error(w, "Error creating room", http.StatusInternalServerError)
			log.Error().Err(err).Msg("Failed to Load booking in database")
			return
		}

		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(byId)
		log.Info().Msg("All bookings by id")
	}
}

func DeleteBooking(client *db.PrismaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]
		deleteBooking, err := client.Booking.FindUnique(
			db.Booking.ID.Equals(id),
		).Delete().Exec(r.Context())

		if err != nil {
			http.Error(w, "Error fetching booking", http.StatusInternalServerError)
			log.Error().Err(err).Msg("Failed to fetch booking from database")
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(deleteBooking)
	}
}

func UpdateBooking(client *db.PrismaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var booking models.BookingModel

		params := mux.Vars(r)
		hotelId := params["hotelId"]
		id := params["id"]

		err := json.NewDecoder(r.Body).Decode(&booking)
		if err != nil {
			http.Error(w, "Invalid Input", http.StatusBadRequest)
			log.Error().Err(err).Msg("Failed to decode request body")
			return
		}

		checkInTime, err := time.Parse("2006-01-02", booking.CheckInDate)
		if err != nil {
			http.Error(w, "Invalid Input", http.StatusBadRequest)
			log.Error().Err(err).Msg("Failed to parse CheckInDate")
			return
		}
		checkOutTime, err := time.Parse("2006-01-02", booking.CheckInDate)
		if err != nil {
			http.Error(w, "Invalid Input", http.StatusBadRequest)
			log.Error().Err(err).Msg("Failed to parse CheckInDate")
			return
		}

		update, err := client.Booking.FindUnique(
			db.Booking.ID.Equals(id),
		).Update(
			db.Booking.CheckInDate.Set(checkInTime),
			db.Booking.CheckOutDate.Set(checkOutTime),
			db.Booking.PaymentStatus.Set(booking.PaymentStatus),
			db.Booking.User.Link(db.User.ID.Equals(booking.UserId)),
			db.Booking.Room.Link(db.Room.ID.Equals(booking.RoomId)),
			db.Booking.Customer.Link(db.Customer.ID.Equals(id)),
			db.Booking.Hotel.Link(db.Hotel.ID.Equals(hotelId)),
			db.Booking.UpdatedAt.Set(time.Now()),
		).Exec(r.Context())
		if err != nil {
			http.Error(w, "Error creating room", http.StatusInternalServerError)
			log.Error().Err(err).Msg("Failed to create booking in database")
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(update)
		log.Info().Msg("Booking has completed")
		fmt.Println(db.Booking.ID)
	}
}
