package repositories

import (
	"encoding/json"
	"os"

	"merchant-bank-api/models"
)

func LoadCustomers() ([]models.Customer, error) {
	var customers []models.Customer
	file, err := os.ReadFile("data/customers.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &customers)
	return customers, err
}

func LoadTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	file, err := os.ReadFile("data/transactions.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &transactions)
	return transactions, err
}

func SaveTransactions(transactions []models.Transaction) error {
	file, err := json.MarshalIndent(transactions, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile("data/transactions.json", file, 0644)
}

func LoadHistory() ([]models.HistoryEntry, error) {
	var history []models.HistoryEntry
	file, err := os.ReadFile("data/history.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &history)
	return history, err
}

func SaveHistory(history []models.HistoryEntry) error {
	file, err := json.MarshalIndent(history, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile("data/history.json", file, 0644)
}