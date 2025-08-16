package auth

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	Pool *pgxpool.Pool
}

type User struct {
	Id         string    `db:"id"`
	Email      string    `db:"email"`
	IsVerified bool      `db:"is_verified"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

var ErrOTPNotFound = errors.New("OTP not found or has expired")
var ErrUserNotFound = errors.New("User not found")

func NewStorage(pool *pgxpool.Pool) *Storage {
	return &Storage{
		Pool: pool,
	}
}

func (s *Storage) SaveOTP(ctx context.Context, email, code string, expiresAt time.Time) error {
	query := `INSERT INTO otp (email, code, expires_at) VALUES ($1, $2, $3)`
	_, err := s.Pool.Exec(ctx, query, email, code, expiresAt)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) FindValidOTP(ctx context.Context, email, code string) (string, error) {
	query := `SELECT id FROM otp WHERE email = $1 AND code = $2 AND expires_at > NOW()`
	var otpId string
	if err := s.Pool.QueryRow(ctx, query, email, code).Scan(&otpId); err != nil {
		if err == pgx.ErrNoRows {
			return "", ErrOTPNotFound
		}
		log.Println(err)
		return "", err
	}
	return otpId, nil
}

func (s *Storage) MarkOTPAsUsed(ctx context.Context, otpId string) error {
	query := `UPDATE otp SET is_used = true WHERE id = $1`
	_, err := s.Pool.Exec(ctx, query, otpId)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) FindUserByEmail(ctx context.Context, email string) (User, error) {
	query := `SELECT * FROM users WHERE email = $1`
	var user User
	if err := s.Pool.QueryRow(ctx, query, email).Scan(&user); err != nil {
		if err == pgx.ErrNoRows {
			return User{}, ErrUserNotFound
		}

		log.Println(err)
		return User{}, err
	}
	return user, nil
}

func (s *Storage) CreateUser(ctx context.Context, email string) (User, error) {
	query := `INSERT INTO users (email, is_verified) VALUES ($1, true) RETURNING *`
	var user User
	if err := s.Pool.QueryRow(ctx, query, email).Scan(&user); err != nil {
		log.Println(err)
		return User{}, err
	}
	return user, nil
}
