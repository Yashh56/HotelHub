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

func CreatePayment(client *db.PrismaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		bookingId := params["bookingId"]

		var payment models.PaymentModel
		err := json.NewDecoder(r.Body).Decode(&payment)
		if err != nil {
			http.Error(w, "Invalid Input", http.StatusBadRequest)
			log.Error().Err(err).Msg("Failed to decode request body")
			return
		}
		paymentDate, err := time.Parse("2006-01-02", payment.PaymentDate)
		if err != nil {
			http.Error(w, "Invalid Input", http.StatusBadRequest)
			log.Error().Err(err).Msg("Failed to parse CheckInDate")
			return
		}
		createPayment, err := client.Payment.CreateOne(
			db.Payment.Amount.Set(payment.Amount),
			db.Payment.PaymentDate.Set(paymentDate),
			db.Payment.PaymentMethod.Set(payment.PaymentMethod),
			db.Payment.Status.Set(payment.Status),
			db.Payment.Booking.Link(db.Booking.ID.Equals(bookingId)),
			db.Payment.CreatedAt.Set(time.Now()),
			db.Payment.UpdatedAt.Set(time.Now()),
		).Exec(r.Context())

		if err != nil {
			http.Error(w, "Error creating room", http.StatusInternalServerError)
			log.Error().Err(err).Msg("Failed to create booking in database")
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createPayment)
		log.Info().Msg("Booking has completed")
		fmt.Println(db.Booking.ID)
	}
}

func DeletePayment(client *db.PrismaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]
		deletePayment, err := client.Payment.FindUnique(
			db.Payment.ID.Equals(id),
		).Delete().Exec(r.Context())

		if err != nil {
			http.Error(w, "Error to delete payment", http.StatusInternalServerError)
			log.Error().Err(err).Msg("Failed to delete payment in database")
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(deletePayment)
	}
}

func PendingPayment(client *db.PrismaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		bookingId := params["bookingId"]

		pending, err := client.Payment.FindMany(
			db.Payment.Booking.Link(db.Booking.ID.Equals(bookingId)),
			db.Payment.Status.Equals("Pending"),
		).Exec(r.Context())

		if err != nil {
			http.Error(w, "Error for getting pending payments", http.StatusInternalServerError)
			log.Error().Err(err).Msg("Failed to retrive pendingful Payments")
			return
		}
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(pending)
	}
}
func SuccessfulPayment(client *db.PrismaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		bookingId := params["bookingId"]

		success, err := client.Payment.FindMany(
			db.Payment.Booking.Link(db.Booking.ID.Equals(bookingId)),
			db.Payment.Status.Equals("Success"),
		).Exec(r.Context())

		if err != nil {
			http.Error(w, "Error for getting success payments", http.StatusInternalServerError)
			log.Error().Err(err).Msg("Failed to retrive Successful Payments")
			return
		}
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(success)
	}
}
