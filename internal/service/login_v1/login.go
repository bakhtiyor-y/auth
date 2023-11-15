package login_v1

import (
	"context"
	"errors"
	"github.com/Shemistan/uzum_auth/internal/models"
	s "github.com/Shemistan/uzum_auth/internal/storage"
	"github.com/Shemistan/uzum_auth/internal/utils/hasher"
	"github.com/Shemistan/uzum_auth/internal/utils/jwt"
	"log"
)

type ILoginService interface {
	Login(ctx context.Context, req *models.AuthUser) (*models.Token, error)
	Check(ctx context.Context) error
}

type serviceLogin struct {
	TokenSecretKey string
	storage        s.IStorage
}

func NewLoginSystemService(TokenSecretKey string, storage s.IStorage) ILoginService {
	return &serviceLogin{
		TokenSecretKey: TokenSecretKey,
		storage:        storage,
	}
}

func (s *serviceLogin) Login(ctx context.Context, req *models.AuthUser) (*models.Token, error) {
	passwordHashOld, err := s.storage.GetPassword(ctx, req.Login)
	if err != nil {
		return nil, err
	}

	if ok := hasher.CheckPassword(passwordHashOld, req.Password); !ok {
		return nil, errors.New("password is not valid")
	}

	res, err := jwt.GenerateTokens(req.Login, "test", s.TokenSecretKey)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *serviceLogin) Check(ctx context.Context) error {
	token, err := jwt.ExtractTokenFromContext(ctx)
	if err != nil {
		return err
	}

	claim, err := jwt.ValidateToken(token, s.TokenSecretKey)
	if err != nil {
		return err
	}

	log.Printf("***<claim>***>>>%+v <<<>>> %s ", claim.Login, claim.Role)

	return nil
}
