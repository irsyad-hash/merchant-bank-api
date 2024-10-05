package utils

import (
	"fmt"
	"time"

	"merchant-bank-api/models"
	"merchant-bank-api/repositories"
)


func GenerateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func LogActivity(action, details string) {
	history, err := repositories.LoadHistory()
	if err != nil {
		fmt.Printf("Error loading history: %v\n", err)
		return
	}

	entry := models.HistoryEntry{
		Timestamp: time.Now().Format(time.RFC3339),
		Action:    action,
		Details:   details,
	}

	history = append(history, entry)
	err = repositories.SaveHistory(history)
	if err != nil {
		fmt.Printf("Error saving history: %v\n", err)
	}
}