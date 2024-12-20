package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"test_baltic/internal/models"
	"test_baltic/internal/services"

	"test_baltic/pkg/utils"
)

var validSourceTypes = map[string]bool{
	"game":    true,
	"server":  true,
	"payment": true,
}

func UserHandler(w http.ResponseWriter, r *http.Request, txService *services.TransactionService) {
	log.Printf("Received %s request at %s", r.Method, r.URL.Path)
	if r.Method == http.MethodPost {
		if len(r.URL.Path) > len("/user/") {
			if r.URL.Path[len(r.URL.Path)-len("/transaction"):] == "/transaction" {
				handleTransaction(w, r, txService)
				return
			}
		}
	} else if r.Method == http.MethodGet {
		handleBalance(w, r, txService)
		return
	}

	utils.WriteError(w, http.StatusNotFound, "Invalid route")
	log.Printf("Invalid route accessed: %s", r.URL.Path)
}

func handleTransaction(w http.ResponseWriter, r *http.Request, txService *services.TransactionService) {
	sourceType := r.Header.Get("Source-Type")
	if !validSourceTypes[sourceType] {
		utils.WriteError(w, http.StatusBadRequest, fmt.Sprintf("Invalid Source-Type: %s", sourceType))
		log.Printf("Invalid Source-Type header: %s", sourceType)
		return
	}

	userID, err := parseUserID(r.URL.Path, "/user/")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid user ID")
		log.Printf("Failed to parse user ID from path %s: %v", r.URL.Path, err)
		return
	}

	var txn models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&txn); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid payload")
		log.Printf("Failed to decode transaction payload for user %d: %v", userID, err)
		return
	}

	log.Printf("Processing transaction for user %d: %+v", userID, txn)

	if txn.State != "win" && txn.State != "lose" {
		utils.WriteError(w, http.StatusBadRequest, "Invalid transaction state")
		log.Printf("Invalid transaction state for user %d: %s", userID, txn.State)
		return
	}

	if _, err := strconv.ParseFloat(txn.Amount, 64); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid amount format")
		log.Printf("Invalid amount format for user %d: %s", userID, txn.Amount)
		return
	}

	err = txService.ProcessTransaction(userID, services.Transaction{
		State:         txn.State,
		Amount:        txn.Amount,
		TransactionID: txn.TransactionID,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to process transaction: %v", err))
		log.Printf("Failed to process transaction for user %d: %v", userID, err)
		return
	}

	log.Printf("Transaction successfully processed for user %d: %+v", userID, txn)
	w.WriteHeader(http.StatusOK)
}

func handleBalance(w http.ResponseWriter, r *http.Request, txService *services.TransactionService) {
	userID, err := parseUserID(r.URL.Path, "/user/")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid user ID")
		log.Printf("Failed to parse user ID from path %s: %v", r.URL.Path, err)
		return
	}

	log.Printf("Retrieving balance for user %d", userID)

	balance, err := txService.GetBalance(userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to retrieve balance: %v", err))
		log.Printf("Failed to retrieve balance for user %d: %v", userID, err)
		return
	}

	response := models.Response{
		UserID:  userID,
		Balance: fmt.Sprintf("%.2f", balance),
	}
	log.Printf("Balance retrieved for user %d: %s", userID, response.Balance)
	utils.WriteJSON(w, http.StatusOK, response)
}

func parseUserID(path, prefix string) (uint64, error) {
	var userID uint64
	_, err := fmt.Sscanf(path, prefix+"%d", &userID)
	return userID, err
}
