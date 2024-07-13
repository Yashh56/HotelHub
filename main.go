package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Yashh56/HotelHub/config"
	"github.com/Yashh56/HotelHub/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load the environment variables !!", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	client, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Could not connect to the database !!", err)
	}

	r := mux.NewRouter()

	routes.UserRoutes(r, client)
	routes.HotelRoutes(r, client)
	routes.RoomRoutes(r, client)
	routes.CustomerRoutes(r, client)
	routes.BookingRoutes(r, client)
	routes.PaymentRoutes(r, client)
	routes.ReviewRoutes(r, client)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://your-frontend-domain.com"}, // Update with your frontend URL
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}).Handler(r)

	server := &http.Server{
		Addr:           ":" + port,
		Handler:        corsHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Starting server on port %s...\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
