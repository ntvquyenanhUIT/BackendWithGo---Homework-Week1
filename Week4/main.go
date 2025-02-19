package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"engineerpro_ex_week4/config"
	"engineerpro_ex_week4/db"
	"engineerpro_ex_week4/handlers"
	"engineerpro_ex_week4/middleware"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file %w", err)
		os.Exit(1)
	}

	// Load config vô struct đã tạo
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config %w", err)
		os.Exit(1)
	}

	// Init database
	dbConn, err := db.InitDB(cfg)
	if err != nil {
		log.Fatal("Error conecting to database %w", err)
		os.Exit(1)
	}

	fmt.Println("Connect Database Successfully")

	defer dbConn.Close()

	userHandler := handlers.NewUserHandler(dbConn)
	authHandler := handlers.NewAuthHanlder(dbConn)

	router := mux.NewRouter()
	router.HandleFunc("/register", authHandler.Register).Methods("POST")
	router.HandleFunc("/login", authHandler.Login).Methods("POST")
	router.HandleFunc("/upload", userHandler.UploadImage).Methods("POST")

	uploadsDir := "./uploads"
	router.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir(uploadsDir))))

	loggedRouter := middleware.LoggingMiddleware(router)
	http.ListenAndServe(":8080", loggedRouter)

}
