package models

type Transaction struct {
	State         string `json:"state"`
	Amount        string `json:"amount"`
	TransactionID string `json:"transactionId"`
}

type Response struct {
	UserID  uint64 `json:"userId"`
	Balance string `json:"balance"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
