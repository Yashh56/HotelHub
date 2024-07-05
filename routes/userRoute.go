package routes

import (
	"github.com/Yashh56/HotelHub/controllers"
	"github.com/Yashh56/HotelHub/prisma/db"
	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router, client *db.PrismaClient) {
	router.HandleFunc("/register", controllers.Register(client)).Methods("POST")
	router.HandleFunc("/login", controllers.Login(client)).Methods("POST")
}
