package controllers

import (
	"encoding/json"
	"net/http"

	"merchant-bank-api/services"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var creds services.LoginCredentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token, err := services.LoginService(creds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Write([]byte(token))
}

func Payment(w http.ResponseWriter, r *http.Request) {
	var transaction services.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = services.PaymentService(transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("username")
	services.LogoutService(username)
	w.WriteHeader(http.StatusOK)
}