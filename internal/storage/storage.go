package storage

import (
	"context"

	"github.com/Shemistan/uzum_auth/internal/models"
)

type IStorage interface {
	CreateUser(ctx context.Context, user *models.CreateUser) error
	MultiCreateUser(ctx context.Context, users []*models.CreateUser) (int64, error)
	UpdateUser(ctx context.Context, user *models.User, login string) error
	DeleteUser(ctx context.Context, login string) error

	GetPassword(ctx context.Context, login string) (string, error)

	ChangePassword(ctx context.Context, req *models.AuthUser) error

	GetUser(ctx context.Context, login string) (*models.User, error)
	GetUsers(ctx context.Context, logins []string) ([]*models.User, error)
	GetAllUsers(ctx context.Context) ([]*models.User, error)
}
