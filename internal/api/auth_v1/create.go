package auth_v1

import (
	"context"

	"github.com/Shemistan/uzum_auth/internal/models"
	pb "github.com/Shemistan/uzum_auth/pkg/auth_v1"
)

func (a *Auth) Create(ctx context.Context, req *pb.Create_Request) (*pb.Create_Response, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	err = a.AuthService.CreateUser(ctx, &models.CreateUser{
		AuthUser: models.AuthUser{
			Login:    req.LoginPassword.Login,
			Password: req.LoginPassword.Password,
		},
		User: a.getModelUser(req.User),
	})

	return &pb.Create_Response{}, err
}

func (a *Auth) getModelUser(req *pb.User) models.User {
	return models.User{
		Role:    req.Role,
		Name:    req.Name,
		Surname: req.Surname,
		Phone:   req.Phone,
		Address: req.Address,
		AddressCoordinate: models.Coordinate{
			X: req.CoordinateAddress.X,
			Y: req.CoordinateAddress.Y,
		},
	}
}
