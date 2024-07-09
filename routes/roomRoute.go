package routes

import (
	"github.com/Yashh56/HotelHub/controllers"
	"github.com/Yashh56/HotelHub/middleware"
	"github.com/Yashh56/HotelHub/prisma/db"
	"github.com/gorilla/mux"
)

func RoomRoutes(router *mux.Router, client *db.PrismaClient) {

	securd := router.PathPrefix("/rooms").Subrouter()
	securd.Use(middleware.AuthMiddleware)

	router.HandleFunc("/rooms/{id}/rooms", controllers.GetRooms(client)).Methods("GET")
	router.HandleFunc("/rooms/{id}/room", controllers.GetRoomById(client)).Methods("GET")

	securd.HandleFunc("/rooms/{id}/create", controllers.CreateRoom(client)).Methods("POST")
	securd.HandleFunc("/rooms/{id}/update", controllers.UpdateRoom(client)).Methods("PUT")
	securd.HandleFunc("/rooms/{id}/delete", controllers.DeleteRoom(client)).Methods("DELETE")
}
