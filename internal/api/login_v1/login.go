package login_v1

import (
	"context"
	"github.com/bakhtiyor-y/auth/internal/models"
	"github.com/bakhtiyor-y/auth/internal/service/login_v1"
	pb "github.com/bakhtiyor-y/auth/pkg/login_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Login struct {
	pb.UnimplementedLoginV1Server
	Service login_v1.ILoginService
}

func (s *Login) Login(ctx context.Context, req *pb.Login_Request) (*pb.Login_Response, error) {

	res, err := s.Service.Login(ctx, &models.AuthUser{
		Login:    req.GetLogin(),
		Password: req.GetPassword(),
	})

	if err != nil {
		return nil, err
	}

	return &pb.Login_Response{
		RefreshToken: res.Refresh,
		AccessToken:  res.Access,
	}, nil
}

func (s *Login) Check(ctx context.Context, _ *pb.Check_Request) (*emptypb.Empty, error) {
	err := s.Service.Check(ctx)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
