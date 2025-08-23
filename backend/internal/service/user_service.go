package service

import (
	"context"
	"fmt"
	"net/mail"
	"time"

	"github.com/elect0/likely/domain"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo  UserRepository
	jwtSecret string
	tokenTTL  time.Duration
}

func NewUserService(repo UserRepository, jwtSecret string, tokenTTL time.Duration) *UserService {
	return &UserService{
		userRepo:  repo,
		jwtSecret: jwtSecret,
		tokenTTL:  tokenTTL,
	}
}

func (s *UserService) SignUp(ctx context.Context, name, email, password string) (*domain.User, string, error) {

	if name == "" {
		return nil, "", fmt.Errorf("name is required")
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return nil, "", fmt.Errorf("invalid email address")
	}

	if len(password) < 8 {
		return nil, "", fmt.Errorf("password must be at least 8 characters long")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, "", fmt.Errorf("failed to hash password: %w", err)
	}

	now := time.Now().UTC()
	user := &domain.User{
		Id:           uuid.New(),
		Name:         name,
		Email:        email,
		PasswordHash: string(passwordHash),
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	createdUser, err := s.userRepo.CreateUser(ctx, user)

	if err != nil {
		return nil, "", fmt.Errorf("failed to create user: %w", err)
	}

	token, err := s.generateJWT(createdUser.Id)
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate token: %w", err)
	}

	return createdUser, token, nil
}

func (s *UserService) SignIn(ctx context.Context, email, password string) (*domain.User, string, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, "", fmt.Errorf("failed to get user by email: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, "", fmt.Errorf("invalid password: %w", err)
	}

	token, err := s.generateJWT(user.Id)

	if err != nil {
		return nil, "", fmt.Errorf("failed to generate token: %w", err)
	}

	return user, token, nil

}

func (s *UserService) generateJWT(userId uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"sub": userId.String(),
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(s.tokenTTL).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}
	return tokenString, nil
}
