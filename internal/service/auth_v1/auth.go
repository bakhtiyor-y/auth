package auth_v1

import (
	"context"
	"github.com/bakhtiyor-y/auth/internal/models"
	repo "github.com/bakhtiyor-y/auth/internal/storage"
	"github.com/bakhtiyor-y/auth/internal/utils/hasher"
	"log"
)

type IAuthSystemService interface {

	CreateUser(ctx context.Context, user *models.CreateUser) error
	MultiCreateUser(ctx context.Context, users []*models.CreateUser) (int64, error)
	UpdateUser(ctx context.Context, user *models.User, login string) error
	DeleteUser(ctx context.Context, login string) error

	ChangePassword(ctx context.Context, req *models.AuthUser) error

	GetUser(ctx context.Context, login string) (*models.User, error)
	GetUsers(ctx context.Context, logins []string) ([]*models.User, error)
	GetAllUsers(ctx context.Context) ([]*models.User, error)
}

type authSystemService struct {
	keyForHashing string
	storage       repo.IStorage
}

func NewAuthSystemService(storage repo.IStorage, keyForHashing string) IAuthSystemService {
	return &authSystemService{
		keyForHashing: keyForHashing,
		storage:       storage,
	}
}

func (a *authSystemService) CreateUser(ctx context.Context, user *models.CreateUser) error {
	passwordHash, err := hasher.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = passwordHash

	return a.storage.CreateUser(ctx, user)
}

func (a *authSystemService) MultiCreateUser(ctx context.Context, users []*models.CreateUser) (int64, error) {
	//TODO implement me
	log.Println("Реализуй меня")
	return 0, nil
}

func (a *authSystemService) UpdateUser(ctx context.Context, user *models.User, login string) error {
	//TODO implement me
	log.Println("Реализуй меня")

	return nil
}

func (a *authSystemService) DeleteUser(ctx context.Context, login string) error {
	//TODO implement me
	log.Println("Реализуй меня")
	return nil
}

func (a *authSystemService) ChangePassword(ctx context.Context, req *models.AuthUser) error {
	//TODO implement me
	log.Println("Реализуй меня")
	return nil
}

func (a *authSystemService) GetUser(ctx context.Context, login string) (*models.User, error) {
	//TODO implement me
	log.Println("Реализуй меня")
	return &models.User{}, nil
}

func (a *authSystemService) GetUsers(ctx context.Context, logins []string) ([]*models.User, error) {
	//TODO implement me
	log.Println("Реализуй меня")
	return nil, nil
}

func (a *authSystemService) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	//TODO implement me
	log.Println("Реализуй меня")
	return nil, nil
}
