package services

import (
	"errors"
	"fmt"
	"time"

	"merchant-bank-api/config"
	"merchant-bank-api/models"
	"merchant-bank-api/repositories"
	"merchant-bank-api/utils"
	"github.com/dgrijalva/jwt-go"
)

type LoginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Transaction struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

func LoginService(creds LoginCredentials) (string, error) {
	customers, err := repositories.LoadCustomers()
	if err != nil {
		return "", err
	}

	for _, customer := range customers {
		if customer.Username == creds.Username && customer.Password == creds.Password {
			config := config.LoadConfig()
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username": creds.Username,
				"exp":      time.Now().Add(time.Hour * 24).Unix(),
			})
			tokenString, err := token.SignedString([]byte(config.JWTSecret))
			if err != nil {
				return "", err
			}
			utils.LogActivity("Login", fmt.Sprintf("User %s logged in", creds.Username))
			return tokenString, nil
		}
	}

	return "", errors.New("Invalid credentials")
}

func PaymentService(transaction Transaction) error {
	customers, err := repositories.LoadCustomers()
	if err != nil {
		return err
	}

	fromExists := false
	toExists := false
	for _, customer := range customers {
		if customer.ID == transaction.From {
			fromExists = true
		}
		if customer.ID == transaction.To {
			toExists = true
		}
	}

	if !fromExists || !toExists {
		return errors.New("Both sender and receiver must be registered customers")
	}

	transactions, err := repositories.LoadTransactions()
	if err != nil {
		return err
	}

	newTransaction := models.Transaction{
		ID:     utils.GenerateID(),
		From:   transaction.From,
		To:     transaction.To,
		Amount: transaction.Amount,
	}

	transactions = append(transactions, newTransaction)
	err = repositories.SaveTransactions(transactions)
	if err != nil {
		return err
	}

	utils.LogActivity("Payment", fmt.Sprintf("Payment of %.2f from %s to %s", transaction.Amount, transaction.From, transaction.To))
	return nil
}

func LogoutService(username string) {
	utils.LogActivity("Logout", fmt.Sprintf("User %s logged out", username))
}