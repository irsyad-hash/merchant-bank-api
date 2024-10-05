package main

import (
	"log"
	"net/http"

	"merchant-bank-api/config"
	"merchant-bank-api/controllers"
	"merchant-bank-api/middlewares"
	"github.com/gorilla/mux"
)

func main() {
	config := config.LoadConfig()
	router := mux.NewRouter()

	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/payment", middlewares.JWTMiddleware(controllers.Payment)).Methods("POST")
	router.HandleFunc("/logout", middlewares.JWTMiddleware(controllers.Logout)).Methods("POST")

	log.Printf("Server starting on port %s\n", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, router))
}
