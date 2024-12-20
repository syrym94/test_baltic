package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"test_baltic/pkg/configs"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // Required for file-based migrations

	"test_baltic/internal/handlers"
	"test_baltic/internal/repos"
	"test_baltic/internal/services"
	db "test_baltic/pkg/db"
)

func main() {
	cfg := configs.LoadConfig()

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode)

	if err := db.InitDB(connStr); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.CloseDB()

	if err := runMigrations(connStr); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	dbInstance := db.GetInstance()
	dbImpl := repos.NewPostgresDB(dbInstance)
	txService := services.NewTransactionService(dbImpl)

	http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		handlers.UserHandler(w, r, txService)
	})

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func runMigrations(connStr string) error {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migration driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///app/migrations",
		"postgres", driver,
	)
	//m, err := migrate.NewWithDatabaseInstance(
	//	"file://migrations",
	//	"postgres", driver,
	//)
	if err != nil {
		return fmt.Errorf("failed to initialize migrations: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	log.Println("Migrations applied successfully")
	return nil
}
