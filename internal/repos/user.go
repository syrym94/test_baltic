package repos

import (
	"database/sql"
	"fmt"
	"test_baltic/internal/models"
)

// Database interface for abstraction
type Database interface {
	GetUserBalance(userID uint64) (float64, error)
	UpdateUserBalance(userID uint64, amount float64, state string, transactionID string) error
	GetUser(userID uint64) (*models.Response, error)
}

type PostgresDB struct {
	DB *sql.DB
}

func NewPostgresDB(db *sql.DB) *PostgresDB {
	return &PostgresDB{DB: db}
}

func (p *PostgresDB) GetUser(userID uint64) (*models.Response, error) {
	var user models.Response
	err := p.DB.QueryRow("SELECT * FROM users WHERE user_id = $1", userID).Scan(&user.UserID, &user.Balance)
	if err != nil {
		return nil, fmt.Errorf("failed to get user %d: %w", userID, err)
	}

	return &user, nil
}

func (p *PostgresDB) GetUserBalance(userID uint64) (float64, error) {
	var balance float64
	err := p.DB.QueryRow("SELECT balance FROM users WHERE user_id = $1", userID).Scan(&balance)
	if err != nil {
		return 0, fmt.Errorf("failed to get balance for user %d: %w", userID, err)
	}

	return balance, nil
}

func (p *PostgresDB) UpdateUserBalance(userID uint64, amount float64, state string, transactionID string) error {
	var balanceUpdate string
	if state == "win" {
		balanceUpdate = "balance + $1"
	} else if state == "lose" {
		balanceUpdate = "balance - $1"
	} else {
		return fmt.Errorf("invalid state: %s", state)
	}

	_, err := p.DB.Exec(`
		WITH updated AS (
			UPDATE users
			SET balance = `+balanceUpdate+`
			WHERE user_id = $2 AND balance + $1 >= 0
			RETURNING user_id
		)
		INSERT INTO transactions (transaction_id, user_id, amount, state)
		SELECT $3, $2, $1, $4
		WHERE EXISTS (SELECT 1 FROM updated);
	`, amount, userID, transactionID, state)

	if err != nil {
		return fmt.Errorf("failed to update balance for user %d: %w", userID, err)
	}

	return nil
}
