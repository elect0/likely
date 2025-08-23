package service

import (
	"context"

	"github.com/elect0/likely/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
}
