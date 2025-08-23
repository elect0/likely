package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/elect0/likely/domain"
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

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func NewPostgresDB(cfg DBConfig) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	db, err := sqlx.Connect("postgres", connStr)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	log.Println("Successfully connected to postgres.")
	return db, nil
}

func (r *PostgresRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	query := "INSERT INTO users (id, name, email, password_hash, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, name, email, created_at, updated_at"

	var createdUser domain.User

	err := r.db.QueryRowxContext(ctx, query, user.Id, user.Name, user.Email, user.PasswordHash, user.CreatedAt, user.UpdatedAt).StructScan(&createdUser)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	fmt.Println(createdUser, "RESPECT")

	return &createdUser, nil
}

func (r *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := "SELECT * FROM users WHERE email = $1"

	var user domain.User

	err := r.db.GetContext(ctx, &user, query, email)

	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	return &user, nil

}
