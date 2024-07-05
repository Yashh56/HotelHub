package routes

import (
	"github.com/Yashh56/HotelHub/controllers"
	"github.com/Yashh56/HotelHub/middleware"
	"github.com/Yashh56/HotelHub/prisma/db"
	"github.com/gorilla/mux"
)

func HotelRoutes(router *mux.Router, client *db.PrismaClient) {

	secured := router.PathPrefix("/hotels").Subrouter()
	secured.Use(middleware.AuthMiddleware)

	// UnProtected Routes
	router.HandleFunc("/hotels", controllers.GetAllHotels(client)).Methods("GET")
	router.HandleFunc("/hotels/{id}", controllers.GetHotelById(client)).Methods("GET")

	// Protected Routes
	secured.HandleFunc("", controllers.CreateHotel(client)).Methods("POST")
	secured.HandleFunc("/{id}", controllers.UpdateHotel(client)).Methods("PUT")
	secured.HandleFunc("/{id}", controllers.DeleteHotel(client)).Methods("DELETE")
}
