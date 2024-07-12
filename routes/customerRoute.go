package routes

import (
	"github.com/Yashh56/HotelHub/controllers"
	"github.com/Yashh56/HotelHub/middleware"
	"github.com/Yashh56/HotelHub/prisma/db"
	"github.com/gorilla/mux"
)

func CustomerRoutes(router *mux.Router, client *db.PrismaClient) {

	securd := router.PathPrefix("/customer").Subrouter()
	securd.Use(middleware.AuthMiddleware)

	securd.HandleFunc("/customer/{hotelId}/all", controllers.GetCustomers(client)).Methods("GET")
	securd.HandleFunc("/customer/{hotelId}/create", controllers.CreateCustomer(client)).Methods("POST")
	securd.HandleFunc("/customer/{id}/delete", controllers.DeleteCustomer(client)).Methods("DELETE")
}
