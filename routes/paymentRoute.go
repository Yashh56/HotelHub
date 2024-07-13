package routes

import (
	"github.com/Yashh56/HotelHub/controllers"
	"github.com/Yashh56/HotelHub/middleware"
	"github.com/Yashh56/HotelHub/prisma/db"
	"github.com/gorilla/mux"
)

func PaymentRoutes(router *mux.Router, client *db.PrismaClient) {

	securd := router.PathPrefix("/payment").Subrouter()
	securd.Use(middleware.AuthMiddleware)

	securd.HandleFunc("/{bookingId}/create", controllers.CreatePayment(client)).Methods("POST")
	securd.HandleFunc("/{id}/delete", controllers.DeletePayment(client)).Methods("POST")
	securd.HandleFunc("/success", controllers.SuccessfulPayment(client)).Methods("GET")
	securd.HandleFunc("/{bookingId}/pending", controllers.PendingPayment(client)).Methods("GET")
}
