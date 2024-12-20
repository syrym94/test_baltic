package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var instance *sql.DB

func InitDB(connStr string) error {
	var err error
	instance, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	return instance.Ping()
}

func GetInstance() *sql.DB {
	return instance
}

func CloseDB() {
	if instance != nil {
		instance.Close()
	}
}
