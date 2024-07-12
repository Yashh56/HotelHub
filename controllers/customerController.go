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

func CreateCustomer(client *db.PrismaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var customer models.CustomerModel

		err := json.NewDecoder(r.Body).Decode(&customer)
		if err != nil {
			http.Error(w, "Invalid Input", http.StatusBadRequest)
			log.Error().Err(err).Msg("Failed to decode request body")
			return
		}

		params := mux.Vars(r)
		id := params["hotelId"]

		newCustomer, err := client.Customer.CreateOne(
			db.Customer.Name.Set(customer.Name),
			db.Customer.Email.Set(customer.Email),
			db.Customer.Phone.Set(customer.Phone),
			db.Customer.Address.Set(customer.Address),
			db.Customer.Hotel.Link(db.Hotel.ID.Equals(id)),
			db.Customer.CreatedAt.Set(time.Now()),
			db.Customer.UpdatedAt.Set(time.Now()),
		).Exec(r.Context())

		if err != nil {
			http.Error(w, "Error Creating customer", http.StatusInternalServerError)
			log.Error().Err(err).Msg("Failed to create a customer in database")
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newCustomer)
		log.Info().Msg("Customer Created Successfully !")
	}
}

func GetCustomers(client *db.PrismaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		id := params["hotelId"]

		allCustomers, err := client.Customer.FindMany(
			db.Customer.Hotel.Link(db.Hotel.ID.Equals(id)),
		).Exec(r.Context())

		if err != nil {
			http.Error(w, "Error getting customer", http.StatusInternalServerError)
			log.Error().Err(err).Msg("Failed to get a customers")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(allCustomers)
	}
}

func DeleteCustomer(client *db.PrismaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		id := params["id"]

		deleteCustomer, err := client.Customer.FindUnique(
			db.Customer.ID.Equals(id),
		).Delete().Exec(r.Context())
		if err != nil {
			http.Error(w, "Error Deleting customer", http.StatusInternalServerError)
			log.Error().Err(err).Msg("Failed to delete a customer in database")
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(deleteCustomer)
		log.Info().Msg("Customer Deleted Successfully !")

	}
}
