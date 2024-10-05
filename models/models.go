package models

type Customer struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Transaction struct {
	ID     string  `json:"id"`
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

type HistoryEntry struct {
	Timestamp string `json:"timestamp"`
	Action    string `json:"action"`
	Details   string `json:"details"`
}