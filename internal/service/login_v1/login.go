package login_v1

import (
	"context"
	"errors"
	"github.com/bakhtiyor-y/auth/internal/models"
	s "github.com/bakhtiyor-y/auth/internal/storage"
	"github.com/bakhtiyor-y/auth/internal/utils/hasher"
	"github.com/bakhtiyor-y/auth/internal/utils/jwt"
	pb "github.com/bakhtiyor-y/auth/pkg/login_v1"
	"log"
	"strconv"
)

type ILoginService interface {
	Login(ctx context.Context, req *models.AuthUser) (*models.Token, error)
	Check(ctx context.Context) (*pb.Check_Response, error)
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

	getUserId, err := s.storage.GetUserId(ctx, req.Login)

	res, err := jwt.GenerateTokens(req.Login, strconv.Itoa(getUserId), s.TokenSecretKey)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *serviceLogin) Check(ctx context.Context) (*pb.Check_Response, error) {

	token, err := jwt.ExtractTokenFromContext(ctx)
	if err != nil {
		return &pb.Check_Response{}, err
	}

	claim, err := jwt.ValidateToken(token, s.TokenSecretKey)
	if err != nil {
		return &pb.Check_Response{}, err
	}

	log.Printf("***<claim>***>>>%+v <<<>>> %s ", claim.Login, claim.UserId)

	return &pb.Check_Response{
		Login:  claim.Login,
		UserId: claim.UserId,
	}, nil
}
