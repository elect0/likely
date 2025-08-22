package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func NewPostgresDB(cfg DBConfig) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	db, err := sqlx.Connect("postgres", connStr)

	if err != nil {
		fmt.Println("EROARE")
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	fmt.Println("PULA")
	log.Println("Successfully connected to postgres.")
	return db, nil
}
