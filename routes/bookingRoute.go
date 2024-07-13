package routes

import (
	"github.com/Yashh56/HotelHub/controllers"
	"github.com/Yashh56/HotelHub/middleware"
	"github.com/Yashh56/HotelHub/prisma/db"
	"github.com/gorilla/mux"
)

func BookingRoutes(router *mux.Router, client *db.PrismaClient) {

	securd := router.PathPrefix("/booking").Subrouter()
	securd.Use(middleware.AuthMiddleware)

	securd.HandleFunc("/{hotelId}/all", controllers.GetBookings(client)).Methods("GET")
	securd.HandleFunc("/{id}/", controllers.GetBookingById(client)).Methods("GET")
	securd.HandleFunc("/{hotelId}/create", controllers.CreateBooking(client)).Methods("POST")
	securd.HandleFunc("/{id}/update", controllers.UpdateBooking(client)).Methods("PUT")
	securd.HandleFunc("/{id}/delete", controllers.DeleteBooking(client)).Methods("DELETE")
}
