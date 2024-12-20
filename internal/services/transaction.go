package services

import (
	"fmt"
	"strconv"
	repo "test_baltic/internal/repos"
)

type Transaction struct {
	State         string
	Amount        string
	TransactionID string
}

type TransactionService struct {
	DB repo.Database
}

func NewTransactionService(db repo.Database) *TransactionService {
	return &TransactionService{DB: db}
}

func (s *TransactionService) ProcessTransaction(userID uint64, txn Transaction) error {
	user, err := s.DB.GetUser(userID)
	if err != nil {
		return err
	}

	if user == nil {
		return fmt.Errorf("user not found")
	}

	if txn.State != "win" && txn.State != "lose" {
		return fmt.Errorf("invalid transaction state: %s", txn.State)
	}

	amount, err := strconv.ParseFloat(txn.Amount, 64)
	if err != nil {
		return fmt.Errorf("invalid amount format: %s", txn.Amount)
	}
	currentBalance, err := strconv.ParseFloat(user.Balance, 64)
	if err != nil {
		return fmt.Errorf("invalid balance format: %s", user.Balance)
	}

	if currentBalance < amount && txn.State == "lose" {
		return fmt.Errorf("user balance is less than lose amount")
	}

	err = s.DB.UpdateUserBalance(userID, amount, txn.State, txn.TransactionID)
	if err != nil {
		return fmt.Errorf("failed to process transaction: %w", err)
	}

	return nil
}

func (s *TransactionService) GetBalance(userID uint64) (float64, error) {
	balance, err := s.DB.GetUserBalance(userID)
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve balance for user %d: %w", userID, err)
	}
	return balance, nil
}
