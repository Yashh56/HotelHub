package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Yashh56/HotelHub/models"
	"github.com/Yashh56/HotelHub/prisma/db"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func CreateReview(client *db.PrismaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["hotelId"]
		userID := r.Context().Value("userId").(string)

		var review models.ReviewModel
		err := json.NewDecoder(r.Body).Decode(&review)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			log.Error().Err(err).Msg("Failed to decode request body")
			return
		}

		createReview, err := client.Review.CreateOne(
			db.Review.Rating.Set(review.Rating),
			db.Review.Comment.Set(review.Comment),
			db.Review.User.Link(db.User.ID.Equals(userID)),
			db.Review.Hotel.Link(db.Hotel.ID.Equals(id)),
		).Exec(r.Context())
		if err != nil {
			http.Error(w, "Error creating review", http.StatusInternalServerError)
			log.Error().Err(err).Msg("Failed to create review in database")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createReview)
		log.Info().Msg("Review created successfully")
	}
}

func GetReviews(client *db.PrismaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		hotelID := vars["hotelId"]

		reviews, err := client.Review.FindMany(
			db.Review.HotelID.Equals(hotelID),
		).Exec(r.Context())
		if err != nil {
			http.Error(w, "Failed to fetch reviews", http.StatusInternalServerError)
			log.Error().Err(err).Msg("Failed to fetch reviews from database")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reviews)
		log.Info().Msg("Reviews fetched successfully")
	}
}
func GetReview(client *db.PrismaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		hotelId := params["hotelId"]
		id := params["id"]
		allReviews, err := client.Review.FindMany(
			db.Review.ID.Equals(id),
			db.Review.Hotel.Link(db.Hotel.ID.Equals(hotelId)),
		).Exec(r.Context())
		if err != nil {
			http.Error(w, "Error creating review", http.StatusInternalServerError)
			log.Error().Err(err).Msg("Failed to get review in database")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(allReviews)

	}
}

func DeleteReview(client *db.PrismaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]

		deleteReview, err := client.Review.FindUnique(
			db.Review.ID.Equals(id),
		).Delete().Exec(r.Context())
		if err != nil {
			http.Error(w, "Error creating review", http.StatusInternalServerError)
			log.Error().Err(err).Msg("Failed to delete review in database")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(deleteReview)
		log.Info().Msg("Review Deleted successfully")
	}
}
