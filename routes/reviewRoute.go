package routes

import (
	"github.com/Yashh56/HotelHub/controllers"
	"github.com/Yashh56/HotelHub/middleware"
	"github.com/Yashh56/HotelHub/prisma/db"
	"github.com/gorilla/mux"
)

func ReviewRoutes(router *mux.Router, client *db.PrismaClient) {

	securd := router.PathPrefix("/review").Subrouter()
	securd.Use(middleware.AuthMiddleware)

	router.HandleFunc("/reviews/{hotelId}/all", controllers.GetReviews(client)).Methods("GET")
	router.HandleFunc("/reviews/{hotelId}/review/{id}", controllers.GetReview(client)).Methods("GET")

	securd.HandleFunc("/review/{hotelId}/create", controllers.CreateReview(client)).Methods("POST")
	securd.HandleFunc("/review/{hotelId}/delete", controllers.DeleteReview(client)).Methods("DELETE")
}
