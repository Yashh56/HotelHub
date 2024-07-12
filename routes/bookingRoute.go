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

	securd.HandleFunc("/booking/{hotelId}/all", controllers.GetBookings(client)).Methods("GET")
	securd.HandleFunc("/booking/{id}/", controllers.GetBookingById(client)).Methods("GET")
	securd.HandleFunc("/booking/{hotelId}/create", controllers.CreateBooking(client)).Methods("POST")
	securd.HandleFunc("/booking/{id}/update", controllers.UpdateBooking(client)).Methods("PUT")
	securd.HandleFunc("/booking/{id}/delete", controllers.DeleteBooking(client)).Methods("DELETE")
}
